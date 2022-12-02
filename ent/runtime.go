// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/ent/schema"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
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
}
