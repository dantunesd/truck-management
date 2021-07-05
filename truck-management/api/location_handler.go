package api

import (
	"net/http"
	"strconv"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locationService *application.LocationService
}

func NewLocationHandler(locationService *application.LocationService) *LocationHandler {
	return &LocationHandler{
		locationService: locationService,
	}
}

func (l *LocationHandler) CreateLocation(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	var input application.CreateLocationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(NewBadRequest(err.Error()))
		return
	}

	result, err := l.locationService.CreateLocation(ID, input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (l *LocationHandler) GetLastLocation(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	result, err := l.locationService.GetLastLocation(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
