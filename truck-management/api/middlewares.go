package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func ErrorHandler(c *gin.Context) {
	c.Next()
	if err := getContextError(c); err != nil {
		response := NewErrorResponse(err)
		c.AbortWithStatusJSON(response.Status, response)
	}
}

func getContextError(c *gin.Context) error {
	errors := c.Errors.ByType(gin.ErrorTypeAny)
	if len(errors) > 0 {
		return errors[0].Err
	}
	return nil
}
