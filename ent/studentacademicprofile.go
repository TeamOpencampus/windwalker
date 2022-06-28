// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"windwalker/ent/studentacademicprofile"
	"windwalker/ent/studentprofile"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
)

// StudentAcademicProfile is the model entity for the StudentAcademicProfile schema.
type StudentAcademicProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Course holds the value of the "course" field.
	Course string `json:"course,omitempty"`
	// Institute holds the value of the "institute" field.
	Institute string `json:"institute,omitempty"`
	// Board holds the value of the "board" field.
	Board string `json:"board,omitempty"`
	// RegNo holds the value of the "reg_no" field.
	RegNo string `json:"reg_no,omitempty"`
	// Department holds the value of the "department" field.
	Department string `json:"department,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate string `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate string `json:"end_date,omitempty"`
	// Marks holds the value of the "marks" field.
	Marks string `json:"marks,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentAcademicProfileQuery when eager-loading is set.
	Edges                            StudentAcademicProfileEdges `json:"edges"`
	student_profile_academic_profile *xid.ID
}

// StudentAcademicProfileEdges holds the relations/edges for other nodes in the graph.
type StudentAcademicProfileEdges struct {
	// StudentProfile holds the value of the student_profile edge.
	StudentProfile *StudentProfile `json:"student_profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StudentProfileOrErr returns the StudentProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentAcademicProfileEdges) StudentProfileOrErr() (*StudentProfile, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*StudentAcademicProfile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case studentacademicprofile.FieldID, studentacademicprofile.FieldCourse, studentacademicprofile.FieldInstitute, studentacademicprofile.FieldBoard, studentacademicprofile.FieldRegNo, studentacademicprofile.FieldDepartment, studentacademicprofile.FieldStartDate, studentacademicprofile.FieldEndDate, studentacademicprofile.FieldMarks:
			values[i] = new(sql.NullString)
		case studentacademicprofile.ForeignKeys[0]: // student_profile_academic_profile
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type StudentAcademicProfile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StudentAcademicProfile fields.
func (sap *StudentAcademicProfile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studentacademicprofile.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sap.ID = value.String
			}
		case studentacademicprofile.FieldCourse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field course", values[i])
			} else if value.Valid {
				sap.Course = value.String
			}
		case studentacademicprofile.FieldInstitute:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field institute", values[i])
			} else if value.Valid {
				sap.Institute = value.String
			}
		case studentacademicprofile.FieldBoard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field board", values[i])
			} else if value.Valid {
				sap.Board = value.String
			}
		case studentacademicprofile.FieldRegNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reg_no", values[i])
			} else if value.Valid {
				sap.RegNo = value.String
			}
		case studentacademicprofile.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field department", values[i])
			} else if value.Valid {
				sap.Department = value.String
			}
		case studentacademicprofile.FieldStartDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				sap.StartDate = value.String
			}
		case studentacademicprofile.FieldEndDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				sap.EndDate = value.String
			}
		case studentacademicprofile.FieldMarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field marks", values[i])
			} else if value.Valid {
				sap.Marks = value.String
			}
		case studentacademicprofile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field student_profile_academic_profile", values[i])
			} else if value.Valid {
				sap.student_profile_academic_profile = new(xid.ID)
				*sap.student_profile_academic_profile = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryStudentProfile queries the "student_profile" edge of the StudentAcademicProfile entity.
func (sap *StudentAcademicProfile) QueryStudentProfile() *StudentProfileQuery {
	return (&StudentAcademicProfileClient{config: sap.config}).QueryStudentProfile(sap)
}

// Update returns a builder for updating this StudentAcademicProfile.
// Note that you need to call StudentAcademicProfile.Unwrap() before calling this method if this StudentAcademicProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (sap *StudentAcademicProfile) Update() *StudentAcademicProfileUpdateOne {
	return (&StudentAcademicProfileClient{config: sap.config}).UpdateOne(sap)
}

// Unwrap unwraps the StudentAcademicProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sap *StudentAcademicProfile) Unwrap() *StudentAcademicProfile {
	tx, ok := sap.config.driver.(*txDriver)
	if !ok {
		panic("ent: StudentAcademicProfile is not a transactional entity")
	}
	sap.config.driver = tx.drv
	return sap
}

// String implements the fmt.Stringer.
func (sap *StudentAcademicProfile) String() string {
	var builder strings.Builder
	builder.WriteString("StudentAcademicProfile(")
	builder.WriteString(fmt.Sprintf("id=%v", sap.ID))
	builder.WriteString(", course=")
	builder.WriteString(sap.Course)
	builder.WriteString(", institute=")
	builder.WriteString(sap.Institute)
	builder.WriteString(", board=")
	builder.WriteString(sap.Board)
	builder.WriteString(", reg_no=")
	builder.WriteString(sap.RegNo)
	builder.WriteString(", department=")
	builder.WriteString(sap.Department)
	builder.WriteString(", start_date=")
	builder.WriteString(sap.StartDate)
	builder.WriteString(", end_date=")
	builder.WriteString(sap.EndDate)
	builder.WriteString(", marks=")
	builder.WriteString(sap.Marks)
	builder.WriteByte(')')
	return builder.String()
}

// StudentAcademicProfiles is a parsable slice of StudentAcademicProfile.
type StudentAcademicProfiles []*StudentAcademicProfile

func (sap StudentAcademicProfiles) config(cfg config) {
	for _i := range sap {
		sap[_i].config = cfg
	}
}
