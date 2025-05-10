package routes

import (
	"time"

	"github.com/Amheklerior/yata/server/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		ExposedHeaders: []string{"Link"},
	}))

	// Set up routes
	r.Get("/healthcheck", app.HealthCheck) // GET /healthcheck
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", app.TasksHandler.HandleGetTasks)       // GET /tasks
		r.Post("/", app.TasksHandler.HandleCreateNewTask) // POST /tasks

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", app.TasksHandler.HandleGetTaskById)   // GET /tasks/:id
			r.Put("/", app.TasksHandler.HandleUpdateTask)    // PUT /tasks/:id
			r.Delete("/", app.TasksHandler.HandleDeleteTask) // DELETE /tasks/:id
		})
	})

	return r
}
