package routes

import (
	"github/SpaceBuckett/CarGo/Internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {

	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/payloads/{id}", app.PayloadHandler.HandleGetPayloadById)

	r.Post("/payloads", app.PayloadHandler.HandleCreatePayload)

	return r
}
