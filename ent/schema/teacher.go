package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
		field.Int32("age_int32"),
		field.Int64("age_int64"),
		field.Uint("age_uint"),
		field.Uint32("age_uint32"),
		field.Uint64("age_uint64"),
		field.Float("weight_float"),
		field.Float32("weight_float32"),
		field.UUID("class_id", uuid.UUID{}),
		field.Time("enroll_at"),
		field.Bool("status"),
		// field.JSON("info", Info{}),
	}
}

//type Info struct {
//	BirthDay time.Time
//	Address  string
//}

func (Teacher) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

// Edges of the Teacher.
func (Teacher) Edges() []ent.Edge {
	return nil
}
