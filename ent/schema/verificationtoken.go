package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VerificationToken holds the schema definition for the VerificationToken entity.
type VerificationToken struct {
	ent.Schema
}

// Fields of the VerificationToken.
func (VerificationToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
	}
}

// Edges of the VerificationToken.
func (VerificationToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required(),
	}
}
