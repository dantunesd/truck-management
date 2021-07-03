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

		c.JSON(http.StatusCreated, result)
		return nil
	}
}

type IDUri struct {
	ID int `uri:"id" binding:"required,numeric"`
}

func (h *TruckHandler) GetHandler() ResponseWrapper {
	return func(c *gin.Context) error {
		var uri IDUri

		if err := c.ShouldBindUri(&uri); err != nil {
			return NewBadRequest(err.Error())
		}

		result, err := h.service.GetTruck(uri.ID)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, result)
		return nil
	}
}

func (h *TruckHandler) DeleteHandler() ResponseWrapper {
	return func(c *gin.Context) error {
		var uri IDUri

		if err := c.ShouldBindUri(&uri); err != nil {
			return NewBadRequest(err.Error())
		}

		if err := h.service.DeleteTruck(uri.ID); err != nil {
			return err
		}

		c.JSON(http.StatusNoContent, nil)
		return nil
	}
}
