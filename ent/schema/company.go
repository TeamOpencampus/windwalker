package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("company_name").NotEmpty(),
		field.String("contact_person_name").NotEmpty(),
		field.String("contact_person_email").NotEmpty(),
		field.String("contact_person_phone").NotEmpty(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("companies").Unique(),
		edge.To("posts", JobPost.Type),
	}
}

// Mixin of the Company.
func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
