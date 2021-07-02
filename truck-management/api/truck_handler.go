package api

import (
	"encoding/json"
	"net/http"
	"truck-management/truck-management/domain"

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

func (h *TruckHandler) CreateHandler() ResponseHandler {
	return func(r *http.Request) (*Response, error) {
		var truck domain.Truck

		if dErr := json.NewDecoder(r.Body).Decode(&truck); dErr != nil {
			return nil, NewBadRequest("invalid content")
		}

		if vErr := validator.New().Struct(truck); vErr != nil {
			return nil, NewBadRequest(vErr.Error())
		}

		result, err := h.service.CreateNewTruck(truck)

		return &Response{result, 201}, err
	}
}

func (h *TruckHandler) GetHandler() ResponseHandler {
	return func(r *http.Request) (*Response, error) {
		var truck domain.Truck
		return &Response{truck, 200}, nil
	}
}
