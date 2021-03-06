// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"windwalker/ent/collegeprofile"
	"windwalker/ent/studentprofile"
	"windwalker/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified bool `json:"email_verified,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// StudentProfile holds the value of the student_profile edge.
	StudentProfile *StudentProfile `json:"student_profile,omitempty"`
	// CollegeProfile holds the value of the college_profile edge.
	CollegeProfile *CollegeProfile `json:"college_profile,omitempty"`
	// Candidates holds the value of the candidates edge.
	Candidates []*User `json:"candidates,omitempty"`
	// EnrolledIn holds the value of the enrolled_in edge.
	EnrolledIn []*User `json:"enrolled_in,omitempty"`
	// Companies holds the value of the companies edge.
	Companies []*Company `json:"companies,omitempty"`
	// JobPosts holds the value of the job_posts edge.
	JobPosts []*JobPost `json:"job_posts,omitempty"`
	// AppliedTo holds the value of the applied_to edge.
	AppliedTo []*JobPost `json:"applied_to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// StudentProfileOrErr returns the StudentProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) StudentProfileOrErr() (*StudentProfile, error) {
	if e.loadedTypes[0] {
		if e.StudentProfile == nil {
			// The edge student_profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: studentprofile.Label}
		}
		return e.StudentProfile, nil
	}
	return nil, &NotLoadedError{edge: "student_profile"}
}

// CollegeProfileOrErr returns the CollegeProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CollegeProfileOrErr() (*CollegeProfile, error) {
	if e.loadedTypes[1] {
		if e.CollegeProfile == nil {
			// The edge college_profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: collegeprofile.Label}
		}
		return e.CollegeProfile, nil
	}
	return nil, &NotLoadedError{edge: "college_profile"}
}

// CandidatesOrErr returns the Candidates value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CandidatesOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Candidates, nil
	}
	return nil, &NotLoadedError{edge: "candidates"}
}

// EnrolledInOrErr returns the EnrolledIn value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EnrolledInOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.EnrolledIn, nil
	}
	return nil, &NotLoadedError{edge: "enrolled_in"}
}

// CompaniesOrErr returns the Companies value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CompaniesOrErr() ([]*Company, error) {
	if e.loadedTypes[4] {
		return e.Companies, nil
	}
	return nil, &NotLoadedError{edge: "companies"}
}

// JobPostsOrErr returns the JobPosts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) JobPostsOrErr() ([]*JobPost, error) {
	if e.loadedTypes[5] {
		return e.JobPosts, nil
	}
	return nil, &NotLoadedError{edge: "job_posts"}
}

// AppliedToOrErr returns the AppliedTo value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AppliedToOrErr() ([]*JobPost, error) {
	if e.loadedTypes[6] {
		return e.AppliedTo, nil
	}
	return nil, &NotLoadedError{edge: "applied_to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmailVerified:
			values[i] = new(sql.NullBool)
		case user.FieldEmail, user.FieldPassword, user.FieldRole:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(xid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified", values[i])
			} else if value.Valid {
				u.EmailVerified = value.Bool
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryStudentProfile queries the "student_profile" edge of the User entity.
func (u *User) QueryStudentProfile() *StudentProfileQuery {
	return (&UserClient{config: u.config}).QueryStudentProfile(u)
}

// QueryCollegeProfile queries the "college_profile" edge of the User entity.
func (u *User) QueryCollegeProfile() *CollegeProfileQuery {
	return (&UserClient{config: u.config}).QueryCollegeProfile(u)
}

// QueryCandidates queries the "candidates" edge of the User entity.
func (u *User) QueryCandidates() *UserQuery {
	return (&UserClient{config: u.config}).QueryCandidates(u)
}

// QueryEnrolledIn queries the "enrolled_in" edge of the User entity.
func (u *User) QueryEnrolledIn() *UserQuery {
	return (&UserClient{config: u.config}).QueryEnrolledIn(u)
}

// QueryCompanies queries the "companies" edge of the User entity.
func (u *User) QueryCompanies() *CompanyQuery {
	return (&UserClient{config: u.config}).QueryCompanies(u)
}

// QueryJobPosts queries the "job_posts" edge of the User entity.
func (u *User) QueryJobPosts() *JobPostQuery {
	return (&UserClient{config: u.config}).QueryJobPosts(u)
}

// QueryAppliedTo queries the "applied_to" edge of the User entity.
func (u *User) QueryAppliedTo() *JobPostQuery {
	return (&UserClient{config: u.config}).QueryAppliedTo(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", email_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailVerified))
	builder.WriteString(", role=")
	builder.WriteString(u.Role)
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
