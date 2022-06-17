package models

type User struct {
	BaseModel
	Email          string `gorm:"unique"`
	PasswordHash   string
	EmailVerified  bool   `gorm:"default:false"`
	Role           string `gorm:"default:user"`
	Profile        Profile
	CollegeProfile CollegeProfile

	// Self Reference
	AssociatedUsers []*User `gorm:"many2many:user_associated_users"`
}
