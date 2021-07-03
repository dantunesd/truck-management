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
	router.POST("/trucks", ErrorHandler(LogHandler(r.truckHandler.CreateHandler(), r.logger)))
	router.GET("/trucks/:id", ErrorHandler(LogHandler(r.truckHandler.GetHandler(), r.logger)))
	router.DELETE("/trucks/:id", ErrorHandler(LogHandler(r.truckHandler.DeleteHandler(), r.logger)))

	return router
}
