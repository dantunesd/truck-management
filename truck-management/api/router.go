package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	gin.SetMode("release")

	router := gin.New()

	router.Use(gin.Recovery())
	router.POST("/trucks", ErrorHandler(LogHandler(rt.truckHandler.CreateHandler(), rt.logger)))
	router.GET("/trucks/:id", ErrorHandler(LogHandler(rt.truckHandler.GetHandler(), rt.logger)))

	return router
}
