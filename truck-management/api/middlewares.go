package api

import (
	"encoding/json"
	"net/http"
	"truck-management/truck-management/domain"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Content    interface{}
	StatusCode int
}

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type ResponseHandler func(r *http.Request) (*Response, error)

func ErrorLogger(logger *logrus.Logger) func(rh ResponseHandler) ResponseHandler {
	return func(rh ResponseHandler) ResponseHandler {
		return func(r *http.Request) (*Response, error) {
			response, err := rh(r)
			if err != nil {
				logger.Error(err)
			}
			return response, err
		}
	}
}

func Responder(rh ResponseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := rh(r)
		if err != nil {
			code, title := getErrorResponse(err)
			responseWriter(w, code, ErrorResponse{title, code})
			return
		}
		responseWriter(w, response.StatusCode, response.Content)
	}
}

func responseWriter(w http.ResponseWriter, code int, content interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	if content != nil {
		body, _ := json.Marshal(content)
		w.Write(body)
	}
}

func getErrorResponse(err error) (int, string) {
	switch terr := err.(type) {
	case *ClientErrors:
		return terr.Code, terr.ErrorMessage
	case *domain.DomainErrors:
		return terr.Code, terr.ErrorMessage
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
