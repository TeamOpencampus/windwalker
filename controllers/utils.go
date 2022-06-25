package controllers

import (
	"errors"
	"net/http"
	"time"
	"windwalker/ent"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Controller struct {
	Client *ent.Client
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(id string, role string) (string, error) {
	claims := &CustomClaims{
		role,
		jwt.RegisteredClaims{
			Subject:   id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret_key"))
}

func NewSuccessResponse[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, data)
}

func RespondEmpty(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func HealthCheck(ctx *gin.Context) {
	NewSuccessResponse(ctx, "all systems operational")
}

func NoRoute(ctx *gin.Context) {
	ctx.Abort()
	_ = ctx.Error(errors.New("router/no-route"))
}

func NoMethod(ctx *gin.Context) {
	ctx.Abort()
	_ = ctx.Error(errors.New("router/no-method"))
}

// Common Handlers

// func GetOnboarding(ctx *gin.Context) {
// 	db := GetDatabase(ctx)
// 	id := ctx.MustGet("ID").(string)
// 	role := ctx.MustGet("ROLE").(string)

// 	var count int64
// 	if role == "user" {
// 		db.Raw("select count(*) from profiles where user_id = ?", id).Scan(&count)
// 	} else {
// 		db.Raw("select count(*) from college_profiles where user_id = ?", id).Scan(&count)
// 	}
// 	if count == 0 {
// 		NewSuccessResponse(ctx, false)
// 	} else {
// 		NewSuccessResponse(ctx, true)
// 	}

// }
