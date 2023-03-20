package schema

import (
	"blog/ent/schema/mixin"

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
		field.String("first_name"),
		field.String("last_name"),
		field.String("email").StructTag(`validate:"email"`),
		field.String("Password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUIDMixin{},
		mixin.TimeMixin{},
	}
}
