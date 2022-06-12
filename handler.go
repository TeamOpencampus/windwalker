package main

import (
	"errors"
	"net/http"
	"time"
	"windwalker/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Server) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "All systems operational.",
	})
}

func (s *Server) Login(ctx *gin.Context) {
	db := GetDatabase(ctx)
	// validate fields
	var data struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	// fetch user
	var user models.User
	if err := db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		msg := "Unknown error occured."
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "Account does not exists."
		}
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   msg,
		})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   "Password does not match.",
		})
		return
	}

	// all ok, return jwt
	sessionToken, err := GetToken(user.ID, user.Role, time.Now().Add(time.Hour*10))
	// refreshToken, err := GetToken(user.ID, time.Now().Add(time.Hour*24*10))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to generate token.",
		})
		return
	}

	ctx.SetCookie("token", sessionToken, 60*60*10, "/", "localhost", false, false)
	// ctx.SetCookie("refresh_token", refreshToken, 60*60*10, "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "Successfully signed in.",
	})
}
func (s *Server) Register(ctx *gin.Context) {
	db := GetDatabase(ctx)
	// validate fields
	var data struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	// check email uniqueness
	var count int64
	if err := db.Model(&models.User{}).Where("email = ?", data.Email).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}
	if count != 0 {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   "An account already exists with this email.",
		})
		return
	}

	// create user
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to hash password.",
		})
		return
	}

	user := models.User{Email: data.Email, PasswordHash: string(hash)}
	tx := db.Create(&user)
	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"data":   "Failed to create account.",
		})
		return
	}

	// all ok, return jwt
	sessionToken, err := GetToken(user.ID, user.Role, time.Now().Add(time.Hour*10))
	// refreshToken, err := GetToken(user.ID, time.Now().Add(time.Hour*24*10))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to generate token.",
		})
		return
	}

	ctx.SetCookie("token", sessionToken, 60*60*10, "/", "localhost", false, false)
	// ctx.SetCookie("refresh_token", refreshToken, 60*60*10, "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "Successfully registered.",
	})
}

func (s *Server) Logout(ctx *gin.Context) {

}

func (s *Server) Verify(ctx *gin.Context) {

}

func (s *Server) Reset(ctx *gin.Context) {

}
