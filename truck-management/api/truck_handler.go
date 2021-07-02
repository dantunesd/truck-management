package api

import (
	"truck-management/truck-management/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ITruckService interface {
	CreateNewTruck(newTruck domain.Truck) (domain.Truck, error)
}

type TruckHandler struct {
	service ITruckService
}

func NewTruckHandler(s ITruckService) *TruckHandler {
	return &TruckHandler{
		service: s,
	}
}

func (h *TruckHandler) CreateHandler() ResponseWrapper {
	return func(c *gin.Context) error {
		var truck domain.Truck

		if err := c.ShouldBindJSON(&truck); err != nil {
			return NewBadRequest("invalid content")
		}

		if vErr := validator.New().Struct(truck); vErr != nil {
			return NewBadRequest(vErr.Error())
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
		c.JSON(200, domain.Truck{})
		return nil
	}
}
