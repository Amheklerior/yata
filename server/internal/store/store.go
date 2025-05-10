package store

import (
	"encoding/json"
	"fmt"
)

// NOTE: I've choosen an integer id over e UUID for (1) project simplicity.
// For larger project I'd carfully try to understand the nature of the data and how they are used,
// and chose integer or UUID accordingly
type TaskId int

// I went with a status field as opposed to a "done" boolean to easily accomodate
// otehr possible values like "blocked" or "in progress"
type TaskStatus string

const (
	TODO      TaskStatus = "todo"
	COMPLETED TaskStatus = "done"
)

// TODO: add audits fields
type Task struct {
	Id     TaskId     `json:"id"`
	Title  string     `json:"title"`
	Detail string     `json:"detail"`
	Status TaskStatus `json:"status"`
}

type TaskStore interface {
	Get() ([]Task, error)
	Create(t *Task) (*Task, error)
	GetById(id TaskId) (*Task, error)
	Update(t *Task) (*Task, error)
	Delete(id TaskId) (bool, error)
}

// a validation func for the status enum (see custom unmarshal func below)
func (ts TaskStatus) IsValid() bool {
	switch ts {
	case TODO, COMPLETED:
		return true
	default:
		return false
	}
}

// NOTE. Go's JSON.Unmarshal (called during decoding ops) doesn't validate enum values by default.
// This custom Unmarshal func add validation for the TaskStatus field, so that it errors if its
// value is not valid.
// This func gets automatically called during any decoding ops involving the TaskStatus type, so
// there's no need to do any manual checking.
func (ts *TaskStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TaskStatus should be a string: %w", err)
	}

	temp := TaskStatus(s)
	if !temp.IsValid() {
		return fmt.Errorf("invalid TaskStatus value: %q", s)
	}

	*ts = temp
	return nil
}
