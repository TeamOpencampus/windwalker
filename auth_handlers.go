package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"windwalker/models"
)

func (s *Server) Login(ctx *gin.Context) {
	db := GetDatabase(ctx)
	// validate fields
	var data struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("auth/invalid-input"))
		return
	}

	// fetch user
	var user models.User
	if err := db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		code := "misc/internal-server-error"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = "auth/account-nonexistent"
		}
		_ = ctx.Error(errors.New(code))
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)); err != nil {
		_ = ctx.Error(errors.New("auth/password-invalid"))
		return
	}

	// all ok, return jwt
	sessionToken, err := GetToken(user.ID, user.Role, time.Now().Add(time.Hour*10))
	// refreshToken, err := GetToken(user.ID, time.Now().Add(time.Hour*24*10))
	if err != nil {
		_ = ctx.Error(errors.New("auth/internal-server-error"))
		return
	}

	ctx.SetCookie("token", sessionToken, 60*60*10, "/", "localhost", false, false)
	// ctx.SetCookie("refresh_token", refreshToken, 60*60*10, "/", "localhost", false, false)
	NewSuccessResponse(ctx, "signed in")
}

func (s *Server) Register(ctx *gin.Context) {
	db := GetDatabase(ctx)
	// validate fields
	var data struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("auth/invalid-input"))
		return
	}

	// check email uniqueness
	var count int64
	if err := db.Model(&models.User{}).Where("email = ?", data.Email).Count(&count).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	if count != 0 {
		_ = ctx.Error(errors.New("auth/email-already-registered"))
		return
	}

	// create user
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	user := models.User{Email: data.Email, PasswordHash: string(hash), Profile: models.Profile{}}
	tx := db.Create(&user)
	if tx.Error != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	// all ok, return jwt
	sessionToken, err := GetToken(user.ID, user.Role, time.Now().Add(time.Hour*10))
	// refreshToken, err := GetToken(user.ID, time.Now().Add(time.Hour*24*10))
	if err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	ctx.SetCookie("token", sessionToken, 60*60*10, "/", "localhost", false, false)
	// ctx.SetCookie("refresh_token", refreshToken, 60*60*10, "/", "localhost", false, false)
	NewSuccessResponse(ctx, "account created")
}

func (s *Server) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", 0, "/", "localhost", false, false)
	NewSuccessResponse(ctx, "logged out")
}

func (s *Server) Reset(ctx *gin.Context) {

}

func (s *Server) VerifyEmail(ctx *gin.Context) {

}
