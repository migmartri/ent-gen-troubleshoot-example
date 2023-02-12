package schema

import (
	"entgo.io/ent"
	_ "github.com/migmartri/test-ent/api/v1"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
