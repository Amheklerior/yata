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
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		http.Error(w, "Invalid path! Missing task id", http.StatusBadRequest)
		return
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task Id! it must be an integer", http.StatusBadRequest)
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
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		http.Error(w, "Invalid path! Missing task id", http.StatusBadRequest)
		return
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task Id! it must be an integer", http.StatusBadRequest)
		return
	}

	// The request body is almost identical to store.Task, but with no id and all fields optional
	var updateTaskReq struct {
		Title  *string           `json:"title"`
		Detail *string           `json:"detail"`
		Status *store.TaskStatus `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateTaskReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fetch existing task and integrate changes...
	existing, err := th.taskStore.GetById(store.TaskId(taskId))
	if err != nil {
		http.Error(w, "Could not update the task", http.StatusInternalServerError)
		return
	}

	if existing == nil {
		http.Error(w, fmt.Sprintf("Could not find task with id %v", taskId), http.StatusNotFound)
		return
	}

	if updateTaskReq.Title != nil {
		existing.Title = *updateTaskReq.Title
	}
	if updateTaskReq.Detail != nil {
		existing.Detail = *updateTaskReq.Detail
	}
	if updateTaskReq.Status != nil {
		existing.Status = *updateTaskReq.Status
	}

	updated, err := th.taskStore.Update(existing)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update task with id %v", int(taskId)), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (th *TasksHandler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		http.Error(w, "Invalid path! Missing task id", http.StatusBadRequest)
		return
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task Id! it must be an integer", http.StatusBadRequest)
		return
	}

	deleted, err := th.taskStore.Delete(store.TaskId(taskId))
	if err != nil {
		http.Error(w, "Failed to delete the task", http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, fmt.Sprintf("Task with id %v does not exists", taskId), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted task with id %d\n", taskId)
}
