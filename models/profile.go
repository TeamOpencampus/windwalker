package models

type BasicProfile struct {
	Name        string `json:"name,omitempty" binding:"required"`
	Phone       string `json:"phone,omitempty" binding:"required"`
	Gender      string `json:"gender,omitempty" binding:"required"`
	Nationality string `json:"nationality,omitempty" binding:"required"`
	Caste       string `json:"caste,omitempty" binding:"required"`
}

type AcademicProfile struct {
	Course     string `json:"course,omitempty" binding:"required"`
	Institute  string `json:"institute,omitempty" binding:"required"`
	Board      string `json:"board,omitempty" binding:"required"`
	RegNo      string `json:"reg_no,omitempty" binding:"required"`
	Department string `json:"department,omitempty" binding:"required"`
	StartDate  string `json:"start_date,omitempty" binding:"required"`
	EndDate    string `json:"end_date,omitempty" binding:"required"`
	Marks      string `json:"marks,omitempty" binding:"required"`
}

type WorkProfile struct {
	Company   string `json:"company,omitempty" binding:"required"`
	StartDate string `json:"start_date,omitempty" binding:"required"`
	EndDate   string `json:"end_date,omitempty" binding:"required"`
	Salary    string `json:"salary,omitempty" binding:"required"`
	Position  string `json:"position,omitempty" binding:"required"`
}

type AdditionalDocuments struct {
	Photo            string `json:"photo,omitempty"`
	Signature        string `json:"signature,omitempty"`
	IdentityProof    string `json:"identity_proof,omitempty"`
	CasteCertificate string `json:"caste_certificate,omitempty"`
}

type CollegeProfile struct {
	Name    string `json:"name,omitempty" binding:"required"`
	Phone   string `json:"phone,omitempty" binding:"required"`
	Address string `json:"address,omitempty" binding:"required"`
	Type    string `json:"type,omitempty" binding:"required"`
}

type Department struct {
	Name     string `json:"name,omitempty"`
	Capacity uint   `json:"capacity,omitempty"`
}
