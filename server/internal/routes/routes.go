package routes

import (
	"github.com/Amheklerior/yata/server/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", app.HealthCheck)

	// NOTE: for a more complex project I'd probably go with routes stacking/grouping
	r.Get("/tasks", app.TasksHandler.HandleGetTasks)
	r.Post("/tasks", app.TasksHandler.HandleCreateNewTask)
	r.Get("/tasks/{id}", app.TasksHandler.HandleGetTaskById)
	r.Put("/tasks/{id}", app.TasksHandler.HandleUpdateTask)
	r.Delete("/tasks/{id}", app.TasksHandler.HandleDeleteTask)

	return r
}
