package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/xid"
)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

//WithAuthentication is a gin middleware which handles JWT sessions.
func WithAuthentication(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.Abort()
			_ = ctx.Error(errors.New("auth/no-header"))
			return
		}
		token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			id, err := xid.FromString(claims.Subject)
			if err != nil {
				ctx.Abort()
				_ = ctx.Error(errors.New("auth/token-parsing-failed"))
				return
			}

			ctx.Set("ID", id)
			ctx.Set("ROLE", claims.Role)
			ctx.Next()
		} else {
			ctx.Abort()
			_ = ctx.Error(errors.New("auth/invalid-token"))
			return
		}
	}
}
