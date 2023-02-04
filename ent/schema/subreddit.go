package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subreddit holds the schema definition for the Subreddit entity.
type Subreddit struct {
	ent.Schema
}

// Fields of the Subreddit.
func (Subreddit) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Comunity name is required").NotEmpty(),
		field.Text("description").Comment("Description is required").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
	}
}

// Edges of the Subreddit.
func (Subreddit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).Ref("subreddit"),
		edge.To("user", User.Type).Required(),
	}
}
