package httpmdw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/pkg/httperrors"
)

func ErrorsMiddleware(errorType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var appError *httperrors.HttpError
		errors := c.Errors.ByType(errorType)
		if len(errors) > 0 {
			err := errors[0].Err
			switch err := err.(type) {
			case *httperrors.HttpError:
				appError = err
			default:
				appError = &httperrors.HttpError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.IndentedJSON(appError.Code, appError)
			c.Abort()
			return
		}
	}
}
