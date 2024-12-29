package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("Student name | 学生姓名"),
		field.Int16("age").
			Comment("Student age | 学生年龄"),
		field.String("address").Optional().
			Comment("Student's home address | 学生家庭住址 "),
		field.Int32("score").Optional().Comment("Student's score | 学生分数"),
		field.Uint32("weight").Optional().Comment("Student's weight | 学生体重"),
		field.Bool("healthy").Optional().Comment("Whether is healthy | 是否健康"),
		field.Int64("code").Optional().Comment("Student's code | 学生编码"),
	}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teachers", Teacher.Type).Ref("students"),
	}
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "students"},
	}
}
