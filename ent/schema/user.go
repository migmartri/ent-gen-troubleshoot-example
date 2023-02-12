package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	v1 "github.com/migmartri/ent-gen-troubleshoot-example/api/v1"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("config").GoType(v1.Test{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
