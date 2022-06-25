// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"windwalker/ent/company"
	"windwalker/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CompanyName holds the value of the "company_name" field.
	CompanyName string `json:"company_name,omitempty"`
	// ContactPersonName holds the value of the "contact_person_name" field.
	ContactPersonName string `json:"contact_person_name,omitempty"`
	// ContactPersonEmail holds the value of the "contact_person_email" field.
	ContactPersonEmail string `json:"contact_person_email,omitempty"`
	// ContactPersonPhone holds the value of the "contact_person_phone" field.
	ContactPersonPhone string `json:"contact_person_phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges          CompanyEdges `json:"edges"`
	user_companies *xid.ID
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*JobPost `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyEdges) UserOrErr() (*User, error) {
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

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) PostsOrErr() ([]*JobPost, error) {
	if e.loadedTypes[1] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldCompanyName, company.FieldContactPersonName, company.FieldContactPersonEmail, company.FieldContactPersonPhone:
			values[i] = new(sql.NullString)
		case company.FieldID:
			values[i] = new(xid.ID)
		case company.ForeignKeys[0]: // user_companies
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Company", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case company.FieldCompanyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company_name", values[i])
			} else if value.Valid {
				c.CompanyName = value.String
			}
		case company.FieldContactPersonName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_person_name", values[i])
			} else if value.Valid {
				c.ContactPersonName = value.String
			}
		case company.FieldContactPersonEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_person_email", values[i])
			} else if value.Valid {
				c.ContactPersonEmail = value.String
			}
		case company.FieldContactPersonPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_person_phone", values[i])
			} else if value.Valid {
				c.ContactPersonPhone = value.String
			}
		case company.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_companies", values[i])
			} else if value.Valid {
				c.user_companies = new(xid.ID)
				*c.user_companies = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Company entity.
func (c *Company) QueryUser() *UserQuery {
	return (&CompanyClient{config: c.config}).QueryUser(c)
}

// QueryPosts queries the "posts" edge of the Company entity.
func (c *Company) QueryPosts() *JobPostQuery {
	return (&CompanyClient{config: c.config}).QueryPosts(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return (&CompanyClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", company_name=")
	builder.WriteString(c.CompanyName)
	builder.WriteString(", contact_person_name=")
	builder.WriteString(c.ContactPersonName)
	builder.WriteString(", contact_person_email=")
	builder.WriteString(c.ContactPersonEmail)
	builder.WriteString(", contact_person_phone=")
	builder.WriteString(c.ContactPersonPhone)
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company

func (c Companies) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
