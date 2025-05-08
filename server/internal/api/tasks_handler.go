package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Amheklerior/yata/server/internal/store"
	"github.com/go-chi/chi/v5"
)

// TODO: add server logs
// TODO: centralize logic for extracting the task id from the url and converting it
// TODO: define utilities to encapsulate the logic for sending a response to the client

type TasksHandler struct {
	taskStore store.TaskStore
}

func NewTasksHandler(taskStore store.TaskStore) *TasksHandler {
	return &TasksHandler{taskStore}
}

func (th *TasksHandler) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.taskStore.Get()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (th *TasksHandler) HandleCreateNewTask(w http.ResponseWriter, r *http.Request) {
	// request body should at least specify the title, optionally the detail, and should not send the id, nor the status
	var createTaskReq struct {
		Title  string  `json:"title"`
		Detail *string `json:"detail"`
	}

	err := json.NewDecoder(r.Body).Decode(&createTaskReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := th.taskStore.Create(&store.Task{
		Id:     -1, // ignored
		Title:  createTaskReq.Title,
		Detail: *createTaskReq.Detail,
		Status: "todo",
	})

	if err != nil {
		http.Error(w, "Failed to create the task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
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

	task, err := th.taskStore.GetById(store.TaskId(taskId))
	if err != nil {
		http.Error(w, "Failed to get the task", http.StatusInternalServerError)
		return
	}

	if task == nil {
		http.Error(w, fmt.Sprintf("Task with id %v does not exist", taskId), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
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
