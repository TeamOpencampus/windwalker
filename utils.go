package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func NewServer(engine *gin.Engine, database *gorm.DB) *Server {
	return &Server{engine: engine, database: database}
}

func WithDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		ctx.Set("DB", db.WithContext(timeoutCtx))
		ctx.Next()
	}
}

//GetDatabase returns database from gin context.
func GetDatabase(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("DB").(*gorm.DB)
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

//WithAuthentication is a gin middleware which handles JWT sessions.
func WithAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"data":   "token cookie is not set.",
			})
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			ctx.Set("ID", claims.ID)
			ctx.Set("ROLE", claims.Role)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"data":   err.Error(),
			})
			return
		}
	}
}

func GetToken(id uint, role string, time time.Time) (string, error) {
	claims := &CustomClaims{
		role,
		jwt.RegisteredClaims{
			ID:        fmt.Sprintf("%d", id),
			ExpiresAt: jwt.NewNumericDate(time),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret_key"))
}

func WithAuthorization(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.MustGet("ROLE").(string) != role {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status": "error",
				"data":   "You don't have permission to view this resource.",
			})
			return
		}
		ctx.Next()
	}
}
