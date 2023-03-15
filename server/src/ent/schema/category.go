package schema

import (
	"blog/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}
