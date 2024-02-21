package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("Teacher's name | 教师姓名"),
		field.Int16("age").
			Annotations(entsql.WithComments(true)).
			Comment("Teacher's age | 教师年龄"),
	}
}

// Edges of the Teacher.
func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("students", Student.Type).Required(),
	}
}

// Mixin of the Teacher.
func (Teacher) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Indexes of the Teacher.
func (Teacher) Indexes() []ent.Index {
	return nil
}

// Annotations of the Teacher
func (Teacher) Annotations() []schema.Annotation {
	return nil
}
