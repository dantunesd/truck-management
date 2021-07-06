package api

import (
	"net/http"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

type ILogger interface {
	Error(args ...interface{})
}

func CreateHandler(
	logger ILogger,
	truckService *application.TruckService,
	locationService *application.LocationService,
	tripService *application.TripService,
) http.Handler {

	truckHandler := NewTruckHandler(truckService)
	locationHandler := NewLocationHandler(locationService)
	tripHandler := NewTripHandler(tripService)

	gin.SetMode("release")
	handler := gin.New()

	handler.Use(gin.Recovery())
	handler.Use(CorsHandler)
	handler.Use(ErrorHandler)
	handler.Use(LogHandler(logger))

	handler.POST("/trucks", truckHandler.CreateTruck)
	handler.GET("/trucks/:id", TruckIdHandler, truckHandler.GetTruck)
	handler.PATCH("/trucks/:id", TruckIdHandler, truckHandler.UpdateTruck)
	handler.DELETE("/trucks/:id", TruckIdHandler, truckHandler.DeleteTruck)

	handler.POST("/trucks/:id/locations", TruckIdHandler, locationHandler.CreateLocation)
	handler.GET("/trucks/:id/locations/last", TruckIdHandler, locationHandler.GetLastLocation)

	handler.GET("/trucks/:id/trips/summary", TruckIdHandler, tripHandler.GetTripSummary)

	return handler
}
