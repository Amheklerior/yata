package api

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Amheklerior/yata/server/internal/store"
	"github.com/go-chi/chi/v5"
)

func TestUpdateTaskHandler(t *testing.T) {

	var testCases = []struct {
		name                string
		id                  string
		body                io.Reader
		expectedStatusCode  int
		expectedUpdatedTask *store.Task
	}{
		{
			name:               "UpdateTaskHandler: success (title update)",
			id:                 "2",
			body:               strings.NewReader(`{"title": "Updated Task Title"}`),
			expectedStatusCode: http.StatusOK,
			expectedUpdatedTask: &store.Task{
				Id:     2,
				Title:  "Updated Task Title",
				Detail: "task with descr",
				Status: "todo",
			},
		},
		{
			name:               "UpdateTaskHandler: success (full update)",
			id:                 "2",
			body:               strings.NewReader(`{"title": "Updated Task Title", "detail": "Updated Task Detail", "status": "done"}`),
			expectedStatusCode: http.StatusOK,
			expectedUpdatedTask: &store.Task{
				Id:     2,
				Title:  "Updated Task Title",
				Detail: "Updated Task Detail",
				Status: "done",
			},
		},
		{
			name:               "UpdateTaskHandler: not found",
			id:                 "5",
			body:               strings.NewReader(`{"title": "Won't matter"}`),
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "UpdateTaskHandler: invalid url param (no id)",
			id:                 "",
			body:               strings.NewReader(`{"title": "Should fail"}`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "UpdateTaskHandler: invalid url param (wrong type for id)",
			id:                 "string-id",
			body:               strings.NewReader(`{"title": "Should also fail"}`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "UpdateTaskHandler: malformed request body (missing quotes)",
			id:                 "2",
			body:               strings.NewReader(`{title: "Missing quotes"}`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "UpdateTaskHandler: malformed request body (invalid field)",
			id:                 "2",
			body:               strings.NewReader(`{"invalidField": "This field does not exist on Task"}`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "UpdateTaskHandler: malformed request body (mispelled field)",
			id:                 "2",
			body:               strings.NewReader(`{"titlee": "It should be 'title'"}`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "UpdateTaskHandler: malformed request body (invalid value)",
			id:                 "2",
			body:               strings.NewReader(`{"status": "completed"}`),
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

			req := httptest.NewRequest(http.MethodGet, "/tasks/{id}", tt.body)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()

			// Act
			handler.HandleUpdateTask(rr, req)

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

			updated, err := testDB.GetById(store.TaskId(id))
			if err != nil {
				t.Fatal("failed to get the updated task")
			}

			if !reflect.DeepEqual(tt.expectedUpdatedTask, updated) {
				t.Errorf("expected tasks %+v; got %+v", tt.expectedUpdatedTask, updated)
			}
		})
	}
}
