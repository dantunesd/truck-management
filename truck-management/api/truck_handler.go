package api

import (
	"net/http"
	"strconv"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"

	"github.com/gin-gonic/gin"
)

type TruckHandler struct {
	service *application.TruckService
}

func NewTruckHandler(s *application.TruckService) *TruckHandler {
	return &TruckHandler{
		service: s,
	}
}

func (h *TruckHandler) CreateTruck(c *gin.Context) {
	var truck domain.Truck

	if err := c.ShouldBindJSON(&truck); err != nil {
		c.Error(NewBadRequest(err.Error()))
		return
	}

	result, err := h.service.CreateNewTruck(truck)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *TruckHandler) GetTruck(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	result, err := h.service.GetTruck(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *TruckHandler) DeleteTruck(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteTruck(ID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
