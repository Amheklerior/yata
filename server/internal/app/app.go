package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
	// TODO: add application data
	// TODO: add handlers
}

func NewApplication() (*Application, error) {
	app := &Application{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "available")
}
