package routes

import (
	"github.com/Amheklerior/yata/server/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", app.HealthCheck)

	// TODO: GET 		/tasks
	// TODO: POST		/tasks
	// TODO: GET 		/tasks/{id}
	// TODO: PUT 		/tasks/{id}
	// TODO: DELETE	/tasks/{id}

	return r
}
