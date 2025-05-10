package api

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Amheklerior/yata/server/internal/store"
)

func TestCreateNewTaskHandler(t *testing.T) {
	var testCases = []struct {
		name                string
		body                io.Reader
		expectedStatusCode  int
		expectedCreatedTask *store.Task
	}{
		{
			name:               "CreateNewTaskHandler: success (title only)",
			body:               strings.NewReader(`{"title": "a valid task"}`),
			expectedStatusCode: http.StatusCreated,
			expectedCreatedTask: &store.Task{
				Id:     1,
				Title:  "a valid task",
				Detail: "",
				Status: "todo",
			},
		},
		{
			name:               "CreateNewTaskHandler: success (title and detail)",
			body:               strings.NewReader(`{"title": "also a valid task", "detail": "with details"}`),
			expectedStatusCode: http.StatusCreated,
			expectedCreatedTask: &store.Task{
				Id:     1,
				Title:  "also a valid task",
				Detail: "with details",
				Status: "todo",
			},
		},
		{
			name:               "CreateNewTaskHandler: malformed request body (missing quotes)",
			body:               strings.NewReader(`{title: "missing quotes"`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "CreateNewTaskHandler: malformed request body (missing mandatory fields)",
			body:               strings.NewReader(`{"detail": "cannot create a task with only detail but no title"`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "CreateNewTaskHandler: malformed request body (invalid values)",
			body:               strings.NewReader(`{"title": 1234}`),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "CreateNewTaskHandler: malformed request body (null title value)",
			body:               strings.NewReader(`{"title": null}`),
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			testDB := store.NewInMemoryTaskStore().With([]store.Task{})
			handler := NewTasksHandler(
				testDB,
				log.New(os.Stdout, "", log.Ldate|log.Ltime),
			)

			req := httptest.NewRequest(http.MethodPost, "/tasks", tt.body)
			rr := httptest.NewRecorder()

			// Act
			handler.HandleCreateNewTask(rr, req)

			// Assert
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("expected status %d; got %d", tt.expectedStatusCode, rr.Code)
			}

			if rr.Code != http.StatusCreated {
				return
			}

			tasks, err := testDB.Get()
			if err != nil {
				t.Fatal("failed to get the task list from testDB")
			}

			created := tasks[0]

			if !reflect.DeepEqual(*tt.expectedCreatedTask, created) {
				t.Errorf("expected tasks %+v; got %+v", *tt.expectedCreatedTask, created)
			}
		})
	}
}
