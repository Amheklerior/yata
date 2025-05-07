package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// TODO: isolate logic for extracting the task id from the url and converting it

type TasksHandler struct{}

func NewTasksHandler() *TasksHandler {
	return &TasksHandler{}
}

func (th *TasksHandler) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	// TODO
	fmt.Fprintln(w, "Got the list of all tasks")
}

func (th *TasksHandler) HandleCreateNewTask(w http.ResponseWriter, r *http.Request) {
	// TODO
	taskId := 1234
	fmt.Fprintf(w, "Created a task with id %d\n", taskId)
}

func (th *TasksHandler) HandleGetTaskById(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "id")
	if taskIdParam == "" {
		http.NotFound(w, r)
		return
	}

	taskId, err := strconv.ParseInt(taskIdParam, 10, 64)
	if err != nil {
		// TODO: implement proper error handling
		http.NotFound(w, r)
		return
	}

	// TODO

	fmt.Fprintf(w, "Got task with id %d\n", taskId)
}

func (th *TasksHandler) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "id")
	if taskIdParam == "" {
		http.NotFound(w, r)
		return
	}

	taskId, err := strconv.ParseInt(taskIdParam, 10, 64)
	if err != nil {
		// TODO: implement proper error handling
		http.NotFound(w, r)
		return
	}

	// TODO

	fmt.Fprintf(w, "Updated task with id %d\n", taskId)
}

func (th *TasksHandler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	taskIdParam := chi.URLParam(r, "id")
	if taskIdParam == "" {
		http.NotFound(w, r)
		return
	}

	taskId, err := strconv.ParseInt(taskIdParam, 10, 64)
	if err != nil {
		// TODO: implement proper error handling
		http.NotFound(w, r)
		return
	}

	// TODO

	fmt.Fprintf(w, "Deleted task with id %d\n", taskId)
}
