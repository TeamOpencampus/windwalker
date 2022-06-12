package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"unique"`
	PasswordHash  string
	EmailVerified bool   `gorm:"default:false"`
	Role          string `gorm:"default:user"`
	Profile       Profile
}
