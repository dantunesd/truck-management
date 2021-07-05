package api

import (
	"net/http"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

type ILogger interface {
	Error(args ...interface{})
}

func CreateHandler(logger ILogger, ts *application.TruckService, ls *application.LocationService, trs *application.TripService) http.Handler {

	truckHandler := NewTruckHandler(ts)
	locationHandler := NewLocationHandler(ls)
	tripHandler := NewTripHandler(trs)

	gin.SetMode("release")
	handler := gin.New()

	handler.Use(gin.Recovery())
	handler.Use(ErrorHandler)
	handler.Use(LogHandler(logger))

	handler.POST("/trucks", truckHandler.CreateTruck)
	handler.GET("/trucks/:id", TruckIdHandler, truckHandler.GetTruck)
	handler.PATCH("/trucks/:id", TruckIdHandler, truckHandler.UpdateTruck)
	handler.DELETE("/trucks/:id", TruckIdHandler, truckHandler.DeleteTruck)

	handler.POST("/trucks/:id/locations", TruckIdHandler, locationHandler.CreateLocation)
	handler.GET("/trucks/:id/locations/last", TruckIdHandler, locationHandler.GetLastLocation)

	handler.GET("/trucks/:id/trips/summary", TruckIdHandler, tripHandler.GetTrip)

	return handler
}
