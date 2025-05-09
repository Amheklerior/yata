package api

import (
	"fmt"
	"net/http"

	"github.com/Amheklerior/yata/server/internal/store"
)

func (th *TasksHandler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := getTaskIdFromURLParam(r)
	if err != nil {
		th.logger.Printf("ERROR: HandleDeleteTask:\n%v\n", err.Error())
		writeJSON(w, http.StatusBadRequest, envelope{"error": err.Error()})
		return
	}

	deleted, err := th.taskStore.Delete(store.TaskId(taskId))
	if err != nil {
		th.logger.Printf("ERROR: HandleDeleteTask: Error deleting task with id %v.\n%v\n", taskId, err.Error())
		writeJSON(w, http.StatusInternalServerError, envelope{"error": "Failed to delete the task"})
		return
	}

	if !deleted {
		th.logger.Printf("INFO: HandleDeleteTask: Task with id %v not deleted. It does not exists.\n", taskId)
		writeJSON(w, http.StatusNotFound, envelope{"error": fmt.Sprintf("Task with id %v does not exists", taskId)})
		return
	}

	th.logger.Printf("INFO: HandleDeleteTask: Successfully deleted task with id %v.\n", taskId)

	writeJSON(w, http.StatusOK, envelope{"message": fmt.Sprintf("Deleted task with id %d\n", taskId)})
}
