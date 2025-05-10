package api

import (
	"encoding/json"
	"net/http"

	"github.com/Amheklerior/yata/server/internal/store"
)

func (th *TasksHandler) HandleCreateNewTask(w http.ResponseWriter, r *http.Request) {
	// request body should at least specify the title, optionally the detail, and should not send the id, nor the status
	var createTaskReq struct {
		Title  string  `json:"title"`
		Detail *string `json:"detail"`
	}

	err := json.NewDecoder(r.Body).Decode(&createTaskReq)
	if err != nil {
		th.logger.Printf("ERROR: HandleCreateNewTask: Error decoding the create task request body.\n%v\n", err.Error())
		writeJSON(w, http.StatusBadRequest, envelope{"error": err.Error()})
		return
	}

	if createTaskReq.Title == "" {
		th.logger.Printf("ERROR: HandleCreateNewTask: Error decoding the request. Empty title.")
		writeJSON(w, http.StatusBadRequest, envelope{"error": "Title cannot be null or empty string"})
		return
	}

	detail := ""
	if createTaskReq.Detail != nil {
		detail = *createTaskReq.Detail
	}

	createdTask, err := th.taskStore.Create(&store.Task{
		Id:     -1, // ignored
		Title:  createTaskReq.Title,
		Detail: detail,
		Status: "todo",
	})

	if err != nil {
		th.logger.Printf("ERROR: HandleCreateNewTask:%v\n", err.Error())
		writeJSON(w, http.StatusInternalServerError, envelope{"error": "Failed to create the task"})
		return
	}

	th.logger.Printf("INFO: HandleCreateNewTask: Successfully created new task.\n%v\n", createdTask)

	writeJSON(w, http.StatusCreated, envelope{"task": createdTask})
}
