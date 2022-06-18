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
	Companies       []Company
	Posts           []*Post `gorm:"many2many:post_students"`
	CompanyID       uint
}
