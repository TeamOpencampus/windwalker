package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func WithAuthorization(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := ctx.MustGet("ROLE").(string)
		if r != role {
			ctx.Abort()
			_ = ctx.Error(errors.New("auth/no-permission"))
			return
		}
		ctx.Next()
	}
}
