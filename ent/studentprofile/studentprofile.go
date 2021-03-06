// Code generated by entc, DO NOT EDIT.

package studentprofile

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the studentprofile type in the database.
	Label = "student_profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldCaste holds the string denoting the caste field in the database.
	FieldCaste = "caste"
	// FieldNationality holds the string denoting the nationality field in the database.
	FieldNationality = "nationality"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeAcademicProfile holds the string denoting the academic_profile edge name in mutations.
	EdgeAcademicProfile = "academic_profile"
	// EdgeWorkProfile holds the string denoting the work_profile edge name in mutations.
	EdgeWorkProfile = "work_profile"
	// Table holds the table name of the studentprofile in the database.
	Table = "student_profiles"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "student_profiles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_student_profile"
	// AcademicProfileTable is the table that holds the academic_profile relation/edge.
	AcademicProfileTable = "student_academic_profiles"
	// AcademicProfileInverseTable is the table name for the StudentAcademicProfile entity.
	// It exists in this package in order to avoid circular dependency with the "studentacademicprofile" package.
	AcademicProfileInverseTable = "student_academic_profiles"
	// AcademicProfileColumn is the table column denoting the academic_profile relation/edge.
	AcademicProfileColumn = "student_profile_academic_profile"
	// WorkProfileTable is the table that holds the work_profile relation/edge.
	WorkProfileTable = "student_work_profiles"
	// WorkProfileInverseTable is the table name for the StudentWorkProfile entity.
	// It exists in this package in order to avoid circular dependency with the "studentworkprofile" package.
	WorkProfileInverseTable = "student_work_profiles"
	// WorkProfileColumn is the table column denoting the work_profile relation/edge.
	WorkProfileColumn = "student_profile_work_profile"
)

// Columns holds all SQL columns for studentprofile fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhone,
	FieldGender,
	FieldCaste,
	FieldNationality,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "student_profiles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_student_profile",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(string) error
	// CasteValidator is a validator for the "caste" field. It is called by the builders before save.
	CasteValidator func(string) error
	// NationalityValidator is a validator for the "nationality" field. It is called by the builders before save.
	NationalityValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)
