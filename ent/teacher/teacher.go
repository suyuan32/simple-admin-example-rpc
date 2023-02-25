// Code generated by ent, DO NOT EDIT.

package teacher

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the teacher type in the database.
	Label = "teacher"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldAgeInt32 holds the string denoting the age_int32 field in the database.
	FieldAgeInt32 = "age_int32"
	// FieldAgeInt64 holds the string denoting the age_int64 field in the database.
	FieldAgeInt64 = "age_int64"
	// FieldAgeUint holds the string denoting the age_uint field in the database.
	FieldAgeUint = "age_uint"
	// FieldAgeUint32 holds the string denoting the age_uint32 field in the database.
	FieldAgeUint32 = "age_uint32"
	// FieldAgeUint64 holds the string denoting the age_uint64 field in the database.
	FieldAgeUint64 = "age_uint64"
	// FieldWeightFloat holds the string denoting the weight_float field in the database.
	FieldWeightFloat = "weight_float"
	// FieldWeightFloat32 holds the string denoting the weight_float32 field in the database.
	FieldWeightFloat32 = "weight_float32"
	// FieldClassID holds the string denoting the class_id field in the database.
	FieldClassID = "class_id"
	// FieldEnrollAt holds the string denoting the enroll_at field in the database.
	FieldEnrollAt = "enroll_at"
	// FieldStatusBool holds the string denoting the status_bool field in the database.
	FieldStatusBool = "status_bool"
	// Table holds the table name of the teacher in the database.
	Table = "teachers"
)

// Columns holds all SQL columns for teacher fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldAge,
	FieldAgeInt32,
	FieldAgeInt64,
	FieldAgeUint,
	FieldAgeUint32,
	FieldAgeUint64,
	FieldWeightFloat,
	FieldWeightFloat32,
	FieldClassID,
	FieldEnrollAt,
	FieldStatusBool,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
