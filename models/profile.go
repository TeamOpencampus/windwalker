package models

type Profile struct {
	BaseModel
	BasicProfile        BasicProfile        `gorm:"embedded" json:"basic_profile,omitempty"`
	AcademicProfile     []AcademicProfile   `json:"academic_profile,omitempty"`
	WorkProfile         []WorkProfile       `json:"work_profile,omitempty"`
	AdditionalDocuments AdditionalDocuments `gorm:"embedded" json:"additional_documents,omitempty"`

	// Reference to User
	UserID uint `json:"-"`
}

type BasicProfile struct {
	Name        string `json:"name,omitempty" binding:"required"`
	Phone       string `json:"phone,omitempty" binding:"required"`
	Gender      string `json:"gender,omitempty" binding:"required"`
	Nationality string `json:"nationality,omitempty" binding:"required"`
	Caste       string `json:"caste,omitempty" binding:"required"`
}

type AcademicProfile struct {
	BaseModel
	Course     string `json:"course,omitempty" binding:"required"`
	Institute  string `json:"institute,omitempty" binding:"required"`
	Board      string `json:"board,omitempty" binding:"required"`
	RegNo      string `json:"reg_no,omitempty" binding:"required"`
	Department string `json:"department,omitempty" binding:"required"`
	StartDate  string `json:"start_date,omitempty" binding:"required"`
	EndDate    string `json:"end_date,omitempty" binding:"required"`
	Marks      string `json:"marks,omitempty" binding:"required"`

	// Reference to Profile
	ProfileID uint `json:"-"`
}

type WorkProfile struct {
	BaseModel
	Company   string `json:"company,omitempty" binding:"required"`
	StartDate string `json:"start_date,omitempty" binding:"required"`
	EndDate   string `json:"end_date,omitempty" binding:"required"`
	Salary    string `json:"salary,omitempty" binding:"required"`
	Position  string `json:"position,omitempty" binding:"required"`

	// Reference to Profile
	ProfileID uint `json:"-"`
}

type AdditionalDocuments struct {
	Photo            string `json:"photo,omitempty"`
	Signature        string `json:"signature,omitempty"`
	IdentityProof    string `json:"identity_proof,omitempty"`
	CasteCertificate string `json:"caste_certificate,omitempty"`
}

type CollegeProfile struct {
	BaseModel
	Name        string       `json:"name,omitempty"`
	Phone       string       `json:"phone,omitempty"`
	Address     string       `json:"address,omitempty"`
	Type        string       `json:"type,omitempty"`
	Departments []Department `json:"departments,omitempty"`

	// Reference
	UserID uint `json:"-"`
}

type Department struct {
	BaseModel
	Name     string `json:"name,omitempty"`
	Capacity uint   `json:"capacity,omitempty"`

	// Reference
	CollegeProfileID uint `json:"-"`
}
