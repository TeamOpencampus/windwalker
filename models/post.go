package models

type Post struct {
	BaseModel
	Title       string `json:"title,omitempty"`
	Location    string `json:"location,omitempty"`
	Salary      string `json:"salary,omitempty"`
	Department  string `json:"department,omitempty"`
	Description string `json:"description,omitempty"`

	// Reference
	CompanyID uint    `json:"-"`
	Students  []*User `gorm:"many2many:post_students" json:"students,omitempty"`
}
