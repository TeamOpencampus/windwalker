// Code generated by entc, DO NOT EDIT.

package company

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCompanyName holds the string denoting the company_name field in the database.
	FieldCompanyName = "company_name"
	// FieldContactPersonName holds the string denoting the contact_person_name field in the database.
	FieldContactPersonName = "contact_person_name"
	// FieldContactPersonEmail holds the string denoting the contact_person_email field in the database.
	FieldContactPersonEmail = "contact_person_email"
	// FieldContactPersonPhone holds the string denoting the contact_person_phone field in the database.
	FieldContactPersonPhone = "contact_person_phone"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "companies"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_companies"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "job_posts"
	// PostsInverseTable is the table name for the JobPost entity.
	// It exists in this package in order to avoid circular dependency with the "jobpost" package.
	PostsInverseTable = "job_posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "company_posts"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldCompanyName,
	FieldContactPersonName,
	FieldContactPersonEmail,
	FieldContactPersonPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "companies"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_companies",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// CompanyNameValidator is a validator for the "company_name" field. It is called by the builders before save.
	CompanyNameValidator func(string) error
	// ContactPersonNameValidator is a validator for the "contact_person_name" field. It is called by the builders before save.
	ContactPersonNameValidator func(string) error
	// ContactPersonEmailValidator is a validator for the "contact_person_email" field. It is called by the builders before save.
	ContactPersonEmailValidator func(string) error
	// ContactPersonPhoneValidator is a validator for the "contact_person_phone" field. It is called by the builders before save.
	ContactPersonPhoneValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)
