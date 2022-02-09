package schema

import (
	"context"
	"fmt"
	"strings"
	"time"

	gen "github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/jonashiltl/sessions-backend/services/user/ent/hook"
	"golang.org/x/crypto/bcrypt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.String("username").
			Unique(),

		field.String("first_name"),

		field.String("last_name").
			Optional(),

		field.String("email").
			Unique(),

		field.String("password"),

		field.String("picture").
			Optional(),

		field.String("blurhash").
			Optional(),

		field.Enum("role").
			Values("ADMIN", "USER").
			Default("USER"),

		field.Time("created_at").
			Default(time.Now),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("id"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// v1.1 should add support for properties on relations
		edge.To("friends", User.Type),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					oldPassword, exists := m.Password()
					if !exists {
						return next.Mutate(ctx, m)
					}

					pBytes, err := bcrypt.GenerateFromPassword([]byte(oldPassword), 10)
					if err != nil {
						return nil, fmt.Errorf("failed to hash password")
					}

					m.SetPassword(string(pBytes))

					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					oldUsername, exists := m.Username()
					if !exists {
						return next.Mutate(ctx, m)
					}

					m.SetUsername(strings.ToLower(oldUsername))

					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
