package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobPost holds the schema definition for the JobPost entity.
type JobPost struct {
	ent.Schema
}

// Fields of the JobPost.
func (JobPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("position").NotEmpty(),
		field.String("location").NotEmpty(),
		field.String("salary").NotEmpty(),
		field.String("description").NotEmpty(),
		field.JSON("tags", []string{}),
		field.Time("created_on").Default(time.Now),
	}
}

// Edges of the JobPost.
func (JobPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("job_posts").Unique(),
		edge.From("company", Company.Type).Ref("posts").Unique(),

		edge.To("candidates", User.Type),
	}
}

// Mixin of the JobPost.
func (JobPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
