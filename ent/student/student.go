// Code generated by ent, DO NOT EDIT.

package student

import (
	"time"
)

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
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
	// FieldAgeInt8 holds the string denoting the age_int8 field in the database.
	FieldAgeInt8 = "age_int8"
	// FieldAgeUint8 holds the string denoting the age_uint8 field in the database.
	FieldAgeUint8 = "age_uint8"
	// FieldAgeInt16 holds the string denoting the age_int16 field in the database.
	FieldAgeInt16 = "age_int16"
	// FieldAgeUint16 holds the string denoting the age_uint16 field in the database.
	FieldAgeUint16 = "age_uint16"
	// FieldAgeInt32 holds the string denoting the age_int32 field in the database.
	FieldAgeInt32 = "age_int32"
	// FieldAgeUint32 holds the string denoting the age_uint32 field in the database.
	FieldAgeUint32 = "age_uint32"
	// FieldAgeInt64 holds the string denoting the age_int64 field in the database.
	FieldAgeInt64 = "age_int64"
	// FieldAgeUint64 holds the string denoting the age_uint64 field in the database.
	FieldAgeUint64 = "age_uint64"
	// FieldAgeInt holds the string denoting the age_int field in the database.
	FieldAgeInt = "age_int"
	// FieldAgeUint holds the string denoting the age_uint field in the database.
	FieldAgeUint = "age_uint"
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
	// Table holds the table name of the student in the database.
	Table = "students"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldAge,
	FieldAgeInt8,
	FieldAgeUint8,
	FieldAgeInt16,
	FieldAgeUint16,
	FieldAgeInt32,
	FieldAgeUint32,
	FieldAgeInt64,
	FieldAgeUint64,
	FieldAgeInt,
	FieldAgeUint,
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
)
