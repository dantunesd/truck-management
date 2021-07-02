package api

import (
	"encoding/json"
	"net/http"
	"truck-management/truck-management/domain"
)

type ErrorResponse struct {
	Title string `json:"title"`
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
	case *domain.DomainErrors:
		return terr.Code
	default:
		return http.StatusInternalServerError
	}
}
