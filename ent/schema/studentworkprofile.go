package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// StudentWorkProfile holds the schema definition for the StudentWorkProfile entity.
type StudentWorkProfile struct {
	ent.Schema
}

// Fields of the StudentWorkProfile.
func (StudentWorkProfile) Fields() []ent.Field {
	return nil
}

// Edges of the StudentWorkProfile.
func (StudentWorkProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student_profile", StudentProfile.Type).Ref("work_profile").Unique(),
	}
}

// Mixin of the StudentWorkProfile.
func (StudentWorkProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
