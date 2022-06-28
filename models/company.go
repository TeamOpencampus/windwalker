package models

type Company struct {
	CompanyName        string `json:"company_name,omitempty" binding:"required"`
	ContactPersonName  string `json:"contact_person_name,omitempty" binding:"required"`
	ContactPersonEmail string `json:"contact_person_email,omitempty" binding:"required"`
	ContactPersonPhone string `json:"contact_person_phone,omitempty" binding:"required"`
}
