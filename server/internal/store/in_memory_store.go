package store

import (
	"errors"
)

type InMemoryTaskStore struct {
	autoIncrementID TaskId
	// NOTE. chose a map over a list as I preferred O(1) lookups by id
	// A list would have allowed O(1) in getting the full list,
	// but I assumed more operations like create/update/delete than getAll.
	tasks map[TaskId]Task
}

func NewInMemoryTaskStore() *InMemoryTaskStore {
	return &InMemoryTaskStore{tasks: make(map[TaskId]Task)}
}

func (s *InMemoryTaskStore) Get() ([]Task, error) {
	// create a list from the map
	list := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		list = append(list, task)
	}

	return list, nil
}

func (s *InMemoryTaskStore) Create(t *Task) (*Task, error) {
	s.autoIncrementID++

	t.Id = s.autoIncrementID
	s.tasks[t.Id] = *t

	return t, nil
}

func (s *InMemoryTaskStore) GetById(id TaskId) (*Task, error) {
	task, ok := s.tasks[id]
	if !ok {
		return nil, nil
	}

	return &task, nil
}

func (s *InMemoryTaskStore) Update(t *Task) (*Task, error) {
	if _, ok := s.tasks[t.Id]; !ok {
		return nil, errors.New("task not found")
	}

	s.tasks[t.Id] = *t
	return t, nil
}

func (s *InMemoryTaskStore) Delete(id TaskId) (bool, error) {
	if _, ok := s.tasks[id]; !ok {
		return false, nil
	}

	delete(s.tasks, id)
	return true, nil
}

// Utility builder function for testing purposes
func (s *InMemoryTaskStore) With(tasks []Task) *InMemoryTaskStore {
	for _, t := range tasks {
		s.tasks[t.Id] = t
	}
	return s
}
