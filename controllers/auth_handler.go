package controllers

import (
	"context"
	"errors"
	"windwalker/ent/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Login(ctx *gin.Context) {
	// db := GetDatabase(ctx)
	// // validate fields
	var data struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("auth/invalid-input"))
		return
	}

	user, err := c.Client.User.
		Query().
		Where(user.Email(data.Email)).
		Only(context.Background())
	if err != nil {
		_ = ctx.Error(errors.New("auth/user-not-found"))
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		_ = ctx.Error(errors.New("auth/invalid-password"))
		return
	}

	// all ok, return jwt
	token, err := GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	NewSuccessResponse(ctx, gin.H{
		"id":    user.ID,
		"role":  user.Role,
		"token": token,
	})
}

func (c *Controller) Register(ctx *gin.Context) {
	// db := GetDatabase(ctx)
	// validate fields
	var data struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("auth/invalid-input"))
		return
	}

	// // create user
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	user, err := c.Client.
		User.
		Create().
		SetEmail(data.Email).
		SetPassword(string(hash)).
		Save(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("auth/email-already-registered"))
		return
	}

	// all ok, return jwt
	token, err := GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	NewSuccessResponse(ctx, gin.H{
		"id":    user.ID,
		"role":  user.Role,
		"token": token,
	})
}

func (c *Controller) Logout(ctx *gin.Context) {
	// NewSuccessResponse(ctx, "logged out")
}

func (c *Controller) Reset(ctx *gin.Context) {

}

func (c *Controller) VerifyEmail(ctx *gin.Context) {

}
