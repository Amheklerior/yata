package store

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
