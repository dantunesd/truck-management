package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILogger interface {
	Error(args ...interface{})
}

type Router struct {
	truckHandler *TruckHandler
	logger       ILogger
}

func NewRouter(handler *TruckHandler, logger ILogger) *Router {
	return &Router{
		truckHandler: handler,
		logger:       logger,
	}
}

func (r Router) GetRoutes() http.Handler {
	gin.SetMode("release")

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(ErrorHandler)
	router.Use(LogHandler(r.logger))

	router.POST("/trucks", r.truckHandler.CreateHandler())
	router.GET("/trucks/:id", TruckIdHandler, r.truckHandler.GetHandler())
	router.DELETE("/trucks/:id", TruckIdHandler, r.truckHandler.DeleteHandler())

	return router
}
