package routes

import (
	"time"

	"github.com/Amheklerior/yata/server/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthcheck", app.HealthCheck)

	// NOTE: for a more complex project I'd probably go with routes stacking/grouping
	r.Get("/tasks", app.TasksHandler.HandleGetTasks)
	r.Post("/tasks", app.TasksHandler.HandleCreateNewTask)
	r.Get("/tasks/{id}", app.TasksHandler.HandleGetTaskById)
	r.Put("/tasks/{id}", app.TasksHandler.HandleUpdateTask)
	r.Delete("/tasks/{id}", app.TasksHandler.HandleDeleteTask)

	return r
}
