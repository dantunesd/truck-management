package api

import (
	"net/http"
	"truck-management/truck-management/domain"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type ResponseWrapper func(c *gin.Context) error

func LogHandler(rw ResponseWrapper, logger ILogger) ResponseWrapper {
	return func(c *gin.Context) error {
		err := rw(c)
		if err != nil {
			logger.Error(err)
		}
		return err
	}
}

func ErrorHandler(rw ResponseWrapper) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := rw(c); err != nil {
			code, title := getErrorResponse(err)
			c.JSON(code, &ErrorResponse{title, code})
		}
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
