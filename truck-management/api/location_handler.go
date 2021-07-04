package api

import (
	"net/http"
	"strconv"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service *application.LocationService
}

func NewLocationHandler(s *application.LocationService) *LocationHandler {
	return &LocationHandler{
		service: s,
	}
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	var location domain.Location

	if err := c.ShouldBindJSON(&location); err != nil {
		c.Error(NewBadRequest(err.Error()))
		return
	}

	result, err := h.service.CreateLocation(ID, location)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *LocationHandler) GetLastLocation(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	result, err := h.service.GetLastLocation(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
