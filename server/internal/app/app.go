package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Amheklerior/yata/server/internal/api"
	"github.com/Amheklerior/yata/server/internal/store"
)

type Application struct {
	Logger       *log.Logger
	TasksHandler *api.TasksHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	store := store.NewInMemoryTaskStore()
	handler := api.NewTasksHandler(store, logger)

	app := &Application{
		Logger:       logger,
		TasksHandler: handler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "available")
}
