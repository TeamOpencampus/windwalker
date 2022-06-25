package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StudentProfile holds the schema definition for the StudentProfile entity.
type StudentProfile struct {
	ent.Schema
}

// Fields of the StudentProfile.
func (StudentProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("gender").NotEmpty(),
		field.String("caste").NotEmpty(),
		field.String("nationality").NotEmpty(),

		// field.String("photo").Default(""),
		// field.String("signature").Default(""),
		// field.String("identity_proof").Default(""),
		// field.String("caste_certificate").Default(""),
	}
}

// Edges of the StudentProfile.
func (StudentProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("student_profile").Unique().Required(),
		edge.To("academic_profile", StudentAcademicProfile.Type),
		edge.To("work_profile", StudentWorkProfile.Type),
	}
}

// Mixin of the StudentProfile.
func (StudentProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
