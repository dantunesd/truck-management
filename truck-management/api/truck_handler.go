package api

import (
	"net/http"
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

func (h *TruckHandler) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func (h *TruckHandler) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uri IDUri

		c.ShouldBindJSON(&uri)

		result, err := h.service.GetTruck(uri.ID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (h *TruckHandler) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uri IDUri

		c.ShouldBindJSON(&uri)

		if err := h.service.DeleteTruck(uri.ID); err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
