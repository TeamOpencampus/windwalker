package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty().StructTag(`json:"-"`),
		field.Bool("email_verified").Default(false),
		field.String("role").Default("student"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student_profile", StudentProfile.Type).Unique(),
		edge.To("college_profile", CollegeProfile.Type).Unique(),
		edge.To("enrolled_in", User.Type).From("candidates"),

		edge.To("companies", Company.Type),
		edge.To("job_posts", JobPost.Type),

		edge.From("applied_to", JobPost.Type).Ref("candidates"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
