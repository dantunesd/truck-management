package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type IDUri struct {
	ID int `uri:"id" binding:"required,numeric"`
}

func TruckIdHandler(c *gin.Context) {
	var uri IDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(400, &ErrorResponse{err.Error(), 400})
		return
	}
	c.Next()
}

func MyLogHandler(logger ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if err := c.Errors.Last(); err != nil {
			logger.Error(err)
		}
	}
}

func MyErrorHandler() gin.HandlerFunc {
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
