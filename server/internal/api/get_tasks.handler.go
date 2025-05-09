package api

import (
	"net/http"
)

func (th *TasksHandler) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.taskStore.Get()
	if err != nil {
		th.logger.Printf("ERROR: HandleGetTasks:\n%v\n", err.Error())
		writeJSON(w, http.StatusInternalServerError, envelope{"error": "Failed to fetch tasks"})
		return
	}

	th.logger.Printf("INFO: HandleGetTasks: Successfully fetched the list of tasks.\n%v\n", tasks)

	writeJSON(w, http.StatusOK, envelope{"tasks": tasks, "total": len(tasks)})
}
