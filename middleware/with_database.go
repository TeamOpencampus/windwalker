package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func WithDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		ctx.Set("DB", db.WithContext(timeoutCtx))
		ctx.Next()
	}
}
