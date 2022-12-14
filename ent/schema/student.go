package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
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
		field.Bool("status_bool"),
		// field.JSON("info", Info{}),
	}
}

//type Info struct {
//	BirthDay time.Time
//	Address  string
//}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
