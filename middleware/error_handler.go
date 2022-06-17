package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(code string, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// handle errors after middleware and routes
		for _, code := range ctx.Errors.Errors() {
			switch code {
			case "auth/no-cookie":
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(code, "cookie `token` not found"))
			default:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(code, "message is not defined"))
			}
			return
		}
	}
}
