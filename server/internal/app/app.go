package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Amheklerior/yata/server/internal/api"
)

type Application struct {
	Logger *log.Logger
	// TODO: add application data
	TasksHandler *api.TasksHandler
}

func NewApplication() (*Application, error) {
	app := &Application{
		Logger:       log.New(os.Stdout, "", log.Ldate|log.Ltime),
		TasksHandler: api.NewTasksHandler(),
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "available")
}
