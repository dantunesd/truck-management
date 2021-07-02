package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	truckHandler *TruckHandler
}

func NewRouter(handler *TruckHandler) *Router {
	return &Router{
		truckHandler: handler,
	}
}

func (rt Router) GetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/trucks", Responser(ErrorLogger(rt.truckHandler.CreateHandler())))
	r.Get("/trucks/{id}", Responser(ErrorLogger(rt.truckHandler.GetHandler())))

	return r
}
