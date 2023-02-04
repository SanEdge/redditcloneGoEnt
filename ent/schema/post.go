package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("postname"),
		field.String("url"),
		field.Text("description"),
		field.Int("voteCount").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required(),
		edge.To("subreddit", Subreddit.Type).Required(),
	}
}
