package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/Amheklerior/yata/server/internal/store"
)

func TestGetTasksHandler(t *testing.T) {
	var testCases = []struct {
		name         string
		initialTasks []store.Task
	}{
		{
			name:         "GetTasksHandler: success (empty list)",
			initialTasks: []store.Task{},
		},
		{
			name: "GetTasksHandler: success (non-empty list)",
			initialTasks: []store.Task{
				{
					Id:     1,
					Title:  "Task #1",
					Detail: "",
					Status: "todo",
				},
				{
					Id:     2,
					Title:  "Task #2",
					Detail: "task with descr",
					Status: "todo",
				},
				{
					Id:     3,
					Title:  "Task #3",
					Detail: "this task is done",
					Status: "done",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			handler := NewTasksHandler(
				store.NewInMemoryTaskStore().With(tt.initialTasks),
				log.New(os.Stdout, "", log.Ldate|log.Ltime),
			)

			req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
			rr := httptest.NewRecorder()

			// Act
			handler.HandleGetTasks(rr, req)

			// Assert
			if rr.Code != http.StatusOK {
				t.Errorf("expected status %d; got %d", http.StatusOK, rr.Code)
			}

			var body struct {
				Tasks []store.Task `json:"tasks"`
				Total int          `json:"total"`
			}
			if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
				t.Fatalf("failed to parse response body: %v", err)
			}

			if body.Total != len(tt.initialTasks) {
				t.Errorf("expected total %d, got %d", len(tt.initialTasks), body.Total)
			}

			if !reflect.DeepEqual(body.Tasks, tt.initialTasks) {
				t.Errorf("expected tasks %+v; got %+v", tt.initialTasks, body.Tasks)
			}
		})
	}
}
