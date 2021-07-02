package api

import (
	"net/http"
	"truck-management/truck-management/application"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router(truckService *application.TruckService) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/trucks", CreateTruckHandler(truckService))
	r.Get("/trucks/{id}", GetTruckHandler())

	return r
}
