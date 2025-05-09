package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Amheklerior/yata/server/internal/store"
)

func (th *TasksHandler) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := getTaskIdFromURLParam(r)
	if err != nil {
		th.logger.Printf("ERROR: HandleUpdateTask:\n%v\n", err.Error())
		writeJSON(w, http.StatusBadRequest, envelope{"error": err.Error()})
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
		writeJSON(w, http.StatusBadRequest, envelope{"error": err.Error()})
		return
	}

	// fetch existing task and integrate changes...
	existing, err := th.taskStore.GetById(store.TaskId(taskId))
	if err != nil {
		th.logger.Printf("ERROR: HandleUpdateTask:\n%v\n", err.Error())
		writeJSON(w, http.StatusInternalServerError, envelope{"error": "Could not update the task"})
		return
	}

	if existing == nil {
		th.logger.Printf("ERROR: HandleUpdateTask: Error while updating task with id %v. It does not exists.\n", taskId)
		writeJSON(w, http.StatusNotFound, envelope{"error": fmt.Sprintf("Could not find task with id %v", taskId)})
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
		writeJSON(w, http.StatusInternalServerError, envelope{"error": fmt.Sprintf("Failed to update task with id %v", int(taskId))})
		return
	}

	th.logger.Printf("INFO: HandleUpdateTask: Successfully updated task with id %v.\n%v\n", taskId, updated)

	writeJSON(w, http.StatusOK, envelope{"task": updated})
}
