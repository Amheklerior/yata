package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Amheklerior/yata/server/internal/store"
	"github.com/Amheklerior/yata/server/internal/utils"
	"github.com/go-chi/chi/v5"
)

// TODO: add server logs
// TODO: centralize logic for extracting the task id from the url and converting it
// TODO: define utilities to encapsulate the logic for sending a response to the client

type TasksHandler struct {
	taskStore store.TaskStore
	logger    *log.Logger
}

func NewTasksHandler(taskStore store.TaskStore, logger *log.Logger) *TasksHandler {
	return &TasksHandler{taskStore, logger}
}

func (th *TasksHandler) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.taskStore.Get()
	if err != nil {
		th.logger.Printf("ERROR: HandleGetTasks:\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to fetch tasks"})
		return
	}

	th.logger.Printf("INFO: HandleGetTasks: Successfully fetched the list of tasks.\n%v\n", tasks)

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"tasks": tasks})
}

func (th *TasksHandler) HandleCreateNewTask(w http.ResponseWriter, r *http.Request) {
	// request body should at least specify the title, optionally the detail, and should not send the id, nor the status
	var createTaskReq struct {
		Title  string  `json:"title"`
		Detail *string `json:"detail"`
	}

	err := json.NewDecoder(r.Body).Decode(&createTaskReq)
	if err != nil {
		th.logger.Printf("ERROR: HandleCreateNewTask: Error decoding the create task request body.\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": err.Error()})
		return
	}

	createdTask, err := th.taskStore.Create(&store.Task{
		Id:     -1, // ignored
		Title:  createTaskReq.Title,
		Detail: *createTaskReq.Detail, // FIXME: potentially referencing a nil pointer
		Status: "todo",
	})

	if err != nil {
		th.logger.Printf("ERROR: HandleCreateNewTask:%v\n", err.Error())
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to create the task"})
		return
	}

	th.logger.Printf("INFO: HandleCreateNewTask: Successfully created new task.\n%v\n", createdTask)

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"task": createdTask})
}

func (th *TasksHandler) HandleGetTaskById(w http.ResponseWriter, r *http.Request) {
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		th.logger.Printf("ERROR: Error extracting the id param from the request url.\n")
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid path! Missing task id"})
		return
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		th.logger.Printf("ERROR: Error parsing the id param from the request url.\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid task Id! it must be an integer"})
		return
	}

	task, err := th.taskStore.GetById(store.TaskId(taskId))
	if err != nil {
		th.logger.Printf("ERROR: HandleGetTaskById:%v\n", err.Error())
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to get the task"})
		return
	}

	if task == nil {
		th.logger.Printf("ERROR: HandleGetTaskById: Error getting task with id %v. It does not exist.\n", taskId)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": fmt.Sprintf("Task with id %v does not exist", taskId)})
		return
	}

	th.logger.Printf("INFO: HandleGetTaskById: Fetched task with id %v\n", taskId)

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"task": task})
}

func (th *TasksHandler) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		th.logger.Printf("ERROR: Error extracting the id param from the request url.\n")
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid path! Missing task id"})
		return
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		th.logger.Printf("ERROR: Error parsing the id param from the request url.\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid task Id! it must be an integer"})
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
		th.logger.Printf("ERROR: HandleUpdateTask: Error decoding request body.\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": err.Error()})
		return
	}

	// fetch existing task and integrate changes...
	existing, err := th.taskStore.GetById(store.TaskId(taskId))
	if err != nil {
		th.logger.Printf("ERROR: HandleUpdateTask:\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Could not update the task"})
		return
	}

	if existing == nil {
		th.logger.Printf("ERROR: HandleUpdateTask: Error while updating task with id %v. It does not exists.\n", taskId)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": fmt.Sprintf("Could not find task with id %v", taskId)})
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

	th.logger.Printf("INFO: HandleUpdateTask: Updating task with id %v...\n%v\n", taskId, existing)

	updated, err := th.taskStore.Update(existing)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": fmt.Sprintf("Failed to update task with id %v", int(taskId))})
		return
	}

	th.logger.Printf("INFO: HandleUpdateTask: Successfully updated task with id %v.\n%v\n", taskId, updated)

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"task": updated})
}

func (th *TasksHandler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		th.logger.Printf("ERROR: Error extracting the id param from the request url.\n")
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid path! Missing task id"})
		return
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		th.logger.Printf("ERROR: Error parsing the id param from the request url.\n%v\n", err.Error())
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid task Id! it must be an integer"})
		return
	}

	deleted, err := th.taskStore.Delete(store.TaskId(taskId))
	if err != nil {
		th.logger.Printf("ERROR: HandleDeleteTask: Error deleting task with id %v.\n%v\n", taskId, err.Error())
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to delete the task"})
		return
	}

	if !deleted {
		th.logger.Printf("INFO: HandleDeleteTask: Task with id %v not deleted. It does not exists.\n", taskId)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": fmt.Sprintf("Task with id %v does not exists", taskId)})
		return
	}

	th.logger.Printf("INFO: HandleDeleteTask: Successfully deleted task with id %v.\n", taskId)

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": fmt.Sprintf("Deleted task with id %d\n", taskId)})
}
