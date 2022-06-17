package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

//GetDatabase returns database from gin context.
func GetDatabase(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("DB").(*gorm.DB)
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GetToken(id uint, role string, time time.Time) (string, error) {
	claims := &CustomClaims{
		role,
		jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", id),
			ExpiresAt: jwt.NewNumericDate(time),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret_key"))
}

type SuccessResponse[T any] struct {
	Data T `json:"data,omitempty"`
}

func NewSuccessResponse[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, SuccessResponse[T]{Data: data})
}
