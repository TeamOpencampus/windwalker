package handlers

import (
	"errors"
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

func NewSuccessResponse[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, data)
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

func GetOnboarding(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	role := ctx.MustGet("ROLE").(string)

	var count int64
	if role == "user" {
		db.Raw("select count(*) from profiles where user_id = ?", id).Scan(&count)
	} else {
		db.Raw("select count(*) from college_profiles where user_id = ?", id).Scan(&count)
	}
	if count == 0 {
		NewSuccessResponse(ctx, false)
	} else {
		NewSuccessResponse(ctx, true)
	}

}
