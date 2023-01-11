// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/suyuan32/simple-admin-example-rpc/ent/schema"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	studentMixin := schema.Student{}.Mixin()
	studentMixinFields0 := studentMixin[0].Fields()
	_ = studentMixinFields0
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescCreatedAt is the schema descriptor for created_at field.
	studentDescCreatedAt := studentMixinFields0[1].Descriptor()
	// student.DefaultCreatedAt holds the default value on creation for the created_at field.
	student.DefaultCreatedAt = studentDescCreatedAt.Default.(func() time.Time)
	// studentDescUpdatedAt is the schema descriptor for updated_at field.
	studentDescUpdatedAt := studentMixinFields0[2].Descriptor()
	// student.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	student.DefaultUpdatedAt = studentDescUpdatedAt.Default.(func() time.Time)
	// student.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	student.UpdateDefaultUpdatedAt = studentDescUpdatedAt.UpdateDefault.(func() time.Time)
	teacherMixin := schema.Teacher{}.Mixin()
	teacherMixinFields0 := teacherMixin[0].Fields()
	_ = teacherMixinFields0
	teacherFields := schema.Teacher{}.Fields()
	_ = teacherFields
	// teacherDescCreatedAt is the schema descriptor for created_at field.
	teacherDescCreatedAt := teacherMixinFields0[1].Descriptor()
	// teacher.DefaultCreatedAt holds the default value on creation for the created_at field.
	teacher.DefaultCreatedAt = teacherDescCreatedAt.Default.(func() time.Time)
	// teacherDescUpdatedAt is the schema descriptor for updated_at field.
	teacherDescUpdatedAt := teacherMixinFields0[2].Descriptor()
	// teacher.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	teacher.DefaultUpdatedAt = teacherDescUpdatedAt.Default.(func() time.Time)
	// teacher.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	teacher.UpdateDefaultUpdatedAt = teacherDescUpdatedAt.UpdateDefault.(func() time.Time)
	// teacherDescID is the schema descriptor for id field.
	teacherDescID := teacherMixinFields0[0].Descriptor()
	// teacher.DefaultID holds the default value on creation for the id field.
	teacher.DefaultID = teacherDescID.Default.(func() uuid.UUID)
}
