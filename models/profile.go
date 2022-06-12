package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	BasicProfile        BasicProfile `gorm:"embedded"`
	AcademicProfile     []AcademicProfile
	WorkProfile         []WorkProfile
	AdditionalDocuments AdditionalDocuments `gorm:"embedded"`

	// Reference to User
	UserID uint
}

type BasicProfile struct {
	Name        string
	Phone       string
	Gender      string
	Nationality string
	Caste       string
}

type AcademicProfile struct {
	gorm.Model
	Course     string
	Institute  string
	Board      string
	RegNo      string
	Department string
	StartDate  string
	EndDate    string
	Marks      string

	// Reference to Profile
	ProfileID uint
}

type WorkProfile struct {
	gorm.Model
	Company   string
	StartDate string
	EndDate   string
	Salary    string
	Position  string

	// Reference to Profile
	ProfileID uint
}

type AdditionalDocuments struct {
	Photo            string
	Signature        string
	IdentityProof    string
	CasteCertificate string
}
