package api

import (
	"net/http"
	"strconv"
	"truck-management/truck-management/application"

	"github.com/gin-gonic/gin"
)

const TRUCK_ID = "id"

type TruckHandler struct {
	truckService *application.TruckService
}

func NewTruckHandler(truckService *application.TruckService) *TruckHandler {
	return &TruckHandler{
		truckService: truckService,
	}
}

func (t *TruckHandler) CreateTruck(c *gin.Context) {
	var truck application.TruckCreateInput

	if err := c.ShouldBindJSON(&truck); err != nil {
		c.Error(NewBadRequest(err.Error()))
		return
	}

	result, err := t.truckService.CreateTruck(truck)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (t *TruckHandler) GetTruck(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(TRUCK_ID))

	result, err := t.truckService.GetTruck(ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (t *TruckHandler) DeleteTruck(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(TRUCK_ID))

	if err := t.truckService.DeleteTruck(ID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (t *TruckHandler) UpdateTruck(c *gin.Context) {
	var truck application.TruckUpdateInput

	if err := c.ShouldBindJSON(&truck); err != nil {
		c.Error(NewBadRequest(err.Error()))
		return
	}

	ID, _ := strconv.Atoi(c.Param(TRUCK_ID))

	err := t.truckService.UpdateTruck(ID, truck)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
