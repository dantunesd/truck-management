package api

import (
	"net/http"
	"strconv"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

type TripHandler struct {
	service *application.TripService
}

func NewTripHandler(s *application.TripService) *TripHandler {
	return &TripHandler{
		service: s,
	}
}

func (h *TripHandler) GetTrip(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	result, err := h.service.GetTrip(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
