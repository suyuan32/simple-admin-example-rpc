// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Comment: "Student name | 学生姓名"},
		{Name: "age", Type: field.TypeInt16, Comment: "Student age | 学生年龄"},
		{Name: "address", Type: field.TypeString, Nullable: true, Comment: "Student's home address | 学生家庭住址 "},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Comment: "Teacher's name | 教师姓名"},
		{Name: "age", Type: field.TypeInt16, Comment: "Teacher's age | 教师年龄"},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
	}
	// TeacherStudentsColumns holds the columns for the "teacher_students" table.
	TeacherStudentsColumns = []*schema.Column{
		{Name: "teacher_id", Type: field.TypeUint64},
		{Name: "student_id", Type: field.TypeUUID},
	}
	// TeacherStudentsTable holds the schema information for the "teacher_students" table.
	TeacherStudentsTable = &schema.Table{
		Name:       "teacher_students",
		Columns:    TeacherStudentsColumns,
		PrimaryKey: []*schema.Column{TeacherStudentsColumns[0], TeacherStudentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "teacher_students_teacher_id",
				Columns:    []*schema.Column{TeacherStudentsColumns[0]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "teacher_students_student_id",
				Columns:    []*schema.Column{TeacherStudentsColumns[1]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StudentsTable,
		TeachersTable,
		TeacherStudentsTable,
	}
)

func init() {
	TeacherStudentsTable.ForeignKeys[0].RefTable = TeachersTable
	TeacherStudentsTable.ForeignKeys[1].RefTable = StudentsTable
}
