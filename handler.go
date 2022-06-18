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

}
