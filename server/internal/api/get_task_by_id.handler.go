package api

import (
	"fmt"
	"net/http"

	"github.com/Amheklerior/yata/server/internal/store"
)

func (th *TasksHandler) HandleGetTaskById(w http.ResponseWriter, r *http.Request) {
	taskId, err := getTaskIdFromURLParam(r)
	if err != nil {
		th.logger.Printf("ERROR: HandleGetTaskById:\n%v\n", err.Error())
		writeJSON(w, http.StatusBadRequest, envelope{"error": err.Error()})
		return
	}

	task, err := th.taskStore.GetById(store.TaskId(taskId))
	if err != nil {
		th.logger.Printf("ERROR: HandleGetTaskById:%v\n", err.Error())
		writeJSON(w, http.StatusInternalServerError, envelope{"error": "Failed to get the task"})
		return
	}

	if task == nil {
		th.logger.Printf("ERROR: HandleGetTaskById: Error getting task with id %v. It does not exist.\n", taskId)
		writeJSON(w, http.StatusNotFound, envelope{"error": fmt.Sprintf("Task with id %v does not exist", taskId)})
		return
	}

	th.logger.Printf("INFO: HandleGetTaskById: Fetched task with id %v\n", taskId)

	writeJSON(w, http.StatusOK, envelope{"task": task})
}
