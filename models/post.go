package models

type Post struct {
	Title       string `json:"title,omitempty"`
	Location    string `json:"location,omitempty"`
	Salary      string `json:"salary,omitempty"`
	Department  string `json:"department,omitempty"`
	Description string `json:"description,omitempty"`
}
