package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

type Router struct {
	truckHandler *TruckHandler
	logger       *logrus.Logger
}

func NewRouter(handler *TruckHandler, logger *logrus.Logger) *Router {
	return &Router{
		truckHandler: handler,
		logger:       logger,
	}
}

func (rt Router) GetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	logger := ErrorLogger(rt.logger)

	r.Post("/trucks", Responder(logger(rt.truckHandler.CreateHandler())))
	r.Get("/trucks/{id}", Responder(logger(rt.truckHandler.GetHandler())))

	return r
}
