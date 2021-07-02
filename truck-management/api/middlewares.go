package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"truck-management/truck-management/domain"
)

type Response struct {
	Content    interface{}
	StatusCode int
}

type ErrorResponse struct {
	Title string `json:"title"`
}

type ResponseHandler func(r *http.Request) (*Response, error)

func ErrorLogger(rh ResponseHandler) ResponseHandler {
	return func(r *http.Request) (*Response, error) {
		response, err := rh(r)
		if err != nil {
			// use a log library
			fmt.Println(err)
		}
		return response, err
	}
}

func Responser(rh ResponseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := rh(r)
		if err != nil {
			responseWriter(w, getHTTPCode(err), ErrorResponse{err.Error()})
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

func getHTTPCode(err error) int {
	switch terr := err.(type) {
	case *ClientErrors:
		return terr.Code
	case *domain.DomainErrors:
		return terr.Code
	default:
		return http.StatusInternalServerError
	}
}
