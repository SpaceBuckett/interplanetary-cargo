package routes

import (
	"github/SpaceBuckett/CarGo/Internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {

	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	return r
}
