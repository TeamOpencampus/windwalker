// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"windwalker/ent/studentprofile"
	"windwalker/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
)

// StudentProfile is the model entity for the StudentProfile schema.
type StudentProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Caste holds the value of the "caste" field.
	Caste string `json:"caste,omitempty"`
	// Nationality holds the value of the "nationality" field.
	Nationality string `json:"nationality,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentProfileQuery when eager-loading is set.
	Edges                StudentProfileEdges `json:"edges"`
	user_student_profile *xid.ID
}

// StudentProfileEdges holds the relations/edges for other nodes in the graph.
type StudentProfileEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// AcademicProfile holds the value of the academic_profile edge.
	AcademicProfile []*StudentAcademicProfile `json:"academic_profile,omitempty"`
	// WorkProfile holds the value of the work_profile edge.
	WorkProfile []*StudentWorkProfile `json:"work_profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentProfileEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// AcademicProfileOrErr returns the AcademicProfile value or an error if the edge
// was not loaded in eager-loading.
func (e StudentProfileEdges) AcademicProfileOrErr() ([]*StudentAcademicProfile, error) {
	if e.loadedTypes[1] {
		return e.AcademicProfile, nil
	}
	return nil, &NotLoadedError{edge: "academic_profile"}
}

// WorkProfileOrErr returns the WorkProfile value or an error if the edge
// was not loaded in eager-loading.
func (e StudentProfileEdges) WorkProfileOrErr() ([]*StudentWorkProfile, error) {
	if e.loadedTypes[2] {
		return e.WorkProfile, nil
	}
	return nil, &NotLoadedError{edge: "work_profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StudentProfile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case studentprofile.FieldName, studentprofile.FieldPhone, studentprofile.FieldGender, studentprofile.FieldCaste, studentprofile.FieldNationality:
			values[i] = new(sql.NullString)
		case studentprofile.FieldID:
			values[i] = new(xid.ID)
		case studentprofile.ForeignKeys[0]: // user_student_profile
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type StudentProfile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StudentProfile fields.
func (sp *StudentProfile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studentprofile.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sp.ID = *value
			}
		case studentprofile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sp.Name = value.String
			}
		case studentprofile.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				sp.Phone = value.String
			}
		case studentprofile.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				sp.Gender = value.String
			}
		case studentprofile.FieldCaste:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field caste", values[i])
			} else if value.Valid {
				sp.Caste = value.String
			}
		case studentprofile.FieldNationality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nationality", values[i])
			} else if value.Valid {
				sp.Nationality = value.String
			}
		case studentprofile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_student_profile", values[i])
			} else if value.Valid {
				sp.user_student_profile = new(xid.ID)
				*sp.user_student_profile = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the StudentProfile entity.
func (sp *StudentProfile) QueryUser() *UserQuery {
	return (&StudentProfileClient{config: sp.config}).QueryUser(sp)
}

// QueryAcademicProfile queries the "academic_profile" edge of the StudentProfile entity.
func (sp *StudentProfile) QueryAcademicProfile() *StudentAcademicProfileQuery {
	return (&StudentProfileClient{config: sp.config}).QueryAcademicProfile(sp)
}

// QueryWorkProfile queries the "work_profile" edge of the StudentProfile entity.
func (sp *StudentProfile) QueryWorkProfile() *StudentWorkProfileQuery {
	return (&StudentProfileClient{config: sp.config}).QueryWorkProfile(sp)
}

// Update returns a builder for updating this StudentProfile.
// Note that you need to call StudentProfile.Unwrap() before calling this method if this StudentProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *StudentProfile) Update() *StudentProfileUpdateOne {
	return (&StudentProfileClient{config: sp.config}).UpdateOne(sp)
}

// Unwrap unwraps the StudentProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *StudentProfile) Unwrap() *StudentProfile {
	tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: StudentProfile is not a transactional entity")
	}
	sp.config.driver = tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *StudentProfile) String() string {
	var builder strings.Builder
	builder.WriteString("StudentProfile(")
	builder.WriteString(fmt.Sprintf("id=%v", sp.ID))
	builder.WriteString(", name=")
	builder.WriteString(sp.Name)
	builder.WriteString(", phone=")
	builder.WriteString(sp.Phone)
	builder.WriteString(", gender=")
	builder.WriteString(sp.Gender)
	builder.WriteString(", caste=")
	builder.WriteString(sp.Caste)
	builder.WriteString(", nationality=")
	builder.WriteString(sp.Nationality)
	builder.WriteByte(')')
	return builder.String()
}

// StudentProfiles is a parsable slice of StudentProfile.
type StudentProfiles []*StudentProfile

func (sp StudentProfiles) config(cfg config) {
	for _i := range sp {
		sp[_i].config = cfg
	}
}