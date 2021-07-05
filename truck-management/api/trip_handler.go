package api

import (
	"net/http"
	"strconv"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

type TripHandler struct {
	tripService *application.TripService
}

func NewTripHandler(tripService *application.TripService) *TripHandler {
	return &TripHandler{
		tripService: tripService,
	}
}

func (t *TripHandler) GetTripSummary(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(TRUCK_ID))

	result, err := t.tripService.GetTrip(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
