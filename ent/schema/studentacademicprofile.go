package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// StudentAcademicProfile holds the schema definition for the StudentAcademicProfile entity.
type StudentAcademicProfile struct {
	ent.Schema
}

// Fields of the StudentAcademicProfile.
func (StudentAcademicProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string { return xid.New().String() }),
		field.String("course").NotEmpty(),
		field.String("institute").NotEmpty(),
		field.String("board").NotEmpty(),
		field.String("reg_no").NotEmpty(),
		field.String("department").Default("No Department"),
		field.String("start_date").NotEmpty(),
		field.String("end_date").NotEmpty(),
		field.String("marks").NotEmpty(),
	}
}

// Edges of the StudentAcademicProfile.
func (StudentAcademicProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student_profile", StudentProfile.Type).Ref("academic_profile").Unique(),
	}
}

// Mixin of the StudentAcademicProfile.
func (StudentAcademicProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
