package api

import (
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

func (h *TruckHandler) CreateHandler() ResponseWrapper {
	return func(c *gin.Context) error {
		var truck domain.Truck

		if err := c.ShouldBindJSON(&truck); err != nil {
			return NewBadRequest(err.Error())
		}

		result, err := h.service.CreateNewTruck(truck)
		if err != nil {
			return err
		}

		c.JSON(201, result)
		return nil
	}
}

func (h *TruckHandler) GetHandler() ResponseWrapper {
	return func(c *gin.Context) error {
		type GetURI struct {
			ID int `uri:"id" binding:"required,numeric"`
		}
		var uri GetURI

		if err := c.ShouldBindUri(&uri); err != nil {
			return NewBadRequest(err.Error())
		}

		result, err := h.service.GetTruck(uri.ID)
		if err != nil {
			return err
		}

		c.JSON(200, result)
		return nil
	}
}
