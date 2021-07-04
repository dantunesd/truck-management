package api

import "net/http"

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

func GetErrorResponse(err error) ErrorResponse {
	switch terr := err.(type) {
	case *ClientErrors:
		return ErrorResponse{terr.ErrorMessage, terr.Code}
	default:
		return ErrorResponse{"internal server error", http.StatusInternalServerError}
	}
}
