package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// CollegeProfile holds the schema definition for the CollegeProfile entity.
type CollegeProfile struct {
	ent.Schema
}

// Fields of the CollegeProfile.
func (CollegeProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string { return xid.New().String() }),
		field.String("name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("address").NotEmpty(),
		field.String("type").NotEmpty(),
	}
}

// Edges of the CollegeProfile.
func (CollegeProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("college_profile").Unique().Required(),
	}
}

// Mixin of the CollegeProfile.
func (CollegeProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
