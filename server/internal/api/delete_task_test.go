package api

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/Amheklerior/yata/server/internal/store"
	"github.com/go-chi/chi/v5"
)

func TestDeleteTaskHandler(t *testing.T) {

	var testCases = []struct {
		name               string
		id                 string
		expectedStatusCode int
	}{
		{
			name:               "DeleteTaskHandler: success",
			id:                 "2",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "DeleteTaskHandler: not found",
			id:                 "5",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "DeleteTaskHandler: invalid url param (no id)",
			id:                 "",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "DeleteTaskHandler: invalid url param (wrong type for id)",
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
			handler.HandleDeleteTask(rr, req)

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

			remainingTests, err := testDB.Get()
			if err != nil {
				t.Fatal("failed to get list of remaining tasks")
			}

			for _, task := range remainingTests {
				if task.Id == store.TaskId(id) {
					t.Errorf("Task with id %v has not been deleted from the database", task.Id)
				}
			}
		})
	}
}
