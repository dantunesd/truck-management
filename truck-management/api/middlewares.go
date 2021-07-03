package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

func TruckIdHandler(c *gin.Context) {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		c.Error(NewBadRequest("id must be numeric"))
		c.Abort()
		return
	}

	c.Next()
}

func LogHandler(logger ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			logger.Error(err)
		}
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(errType gin.ErrorType) gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Next()

			errors := c.Errors.ByType(errType)
			if len(errors) > 0 {
				err := errors[0].Err
				code, title := GetErrorResponse(err)
				c.AbortWithStatusJSON(code, &ErrorResponse{title, code})
			}
		}
	}(gin.ErrorTypeAny)
}

func GetErrorResponse(err error) (int, string) {
	switch terr := err.(type) {
	case *ClientErrors:
		return terr.Code, terr.ErrorMessage
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
