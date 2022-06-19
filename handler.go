package main

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) HealthCheck(ctx *gin.Context) {
	NewSuccessResponse(ctx, "all systems operational")
}

func (s *Server) NoRoute(ctx *gin.Context) {
	ctx.Abort()
	_ = ctx.Error(errors.New("router/no-route"))
}

func (s *Server) NoMethod(ctx *gin.Context) {
	ctx.Abort()
	_ = ctx.Error(errors.New("router/no-method"))
}

// Common Handlers

func (s *Server) GetOnboarding(ctx *gin.Context) {
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
