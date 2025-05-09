package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/Amheklerior/yata/server/internal/store"
	"github.com/go-chi/chi/v5"
)

func TestGetTaskByIdHandler(t *testing.T) {

	var testCases = []struct {
		name               string
		id                 string
		expectedStatusCode int
	}{
		{
			name:               "GetTaskByIdHandler: success",
			id:                 "2",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "GetTaskByIdHandler: not found",
			id:                 "5",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "GetTaskByIdHandler: invalid url param (no id)",
			id:                 "",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "GetTaskByIdHandler: invalid url param (wrong type for id)",
			id:                 "string-id",
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	// test db setup
	tasks := []store.Task{
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
	}
	testDB := store.NewInMemoryTaskStore().With(tasks)

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			handler := NewTasksHandler(
				testDB,
				log.New(os.Stdout, "", log.Ldate|log.Ltime),
			)

			req := httptest.NewRequest(http.MethodGet, "/tasks/{id}", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()

			// Act
			handler.HandleGetTaskById(rr, req)

			// Assert
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("expected status %d; got %d", tt.expectedStatusCode, rr.Code)
			}

			if rr.Code != http.StatusOK {
				return
			}

			id, err := strconv.Atoi(tt.id)
			if err != nil {
				t.Fatalf("failed to convert id %v", tt.id)
			}

			var body struct {
				Task store.Task `json:"task"`
			}

			if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
				t.Fatalf("failed to parse response body: %v", err)
			}

			for _, task := range tasks {
				if task.Id == store.TaskId(id) && !reflect.DeepEqual(body.Task, task) {
					t.Errorf("expected tasks %+v; got %+v", task, body.Task)
					return
				}
			}
		})
	}
}
