package api

import (
	"net/http"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

type ILogger interface {
	Error(args ...interface{})
}

func CreateHandler(logger ILogger, ts *application.TruckService) http.Handler {

	truckHandler := NewTruckHandler(ts)

	gin.SetMode("release")
	handler := gin.New()

	handler.Use(gin.Recovery())
	handler.Use(ErrorHandler)
	handler.Use(LogHandler(logger))

	handler.POST("/trucks", truckHandler.CreateTruck)
	handler.GET("/trucks/:id", TruckIdHandler, truckHandler.GetTruck)
	handler.DELETE("/trucks/:id", TruckIdHandler, truckHandler.DeleteTruck)

	return handler
}
