// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StudentUpdate) SetUpdatedAt(t time.Time) *StudentUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetStatus sets the "status" field.
func (su *StudentUpdate) SetStatus(u uint8) *StudentUpdate {
	su.mutation.ResetStatus()
	su.mutation.SetStatus(u)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *StudentUpdate) SetNillableStatus(u *uint8) *StudentUpdate {
	if u != nil {
		su.SetStatus(*u)
	}
	return su
}

// AddStatus adds u to the "status" field.
func (su *StudentUpdate) AddStatus(u int8) *StudentUpdate {
	su.mutation.AddStatus(u)
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *StudentUpdate) ClearStatus() *StudentUpdate {
	su.mutation.ClearStatus()
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *StudentUpdate) SetNillableName(s *string) *StudentUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetAge sets the "age" field.
func (su *StudentUpdate) SetAge(i int16) *StudentUpdate {
	su.mutation.ResetAge()
	su.mutation.SetAge(i)
	return su
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (su *StudentUpdate) SetNillableAge(i *int16) *StudentUpdate {
	if i != nil {
		su.SetAge(*i)
	}
	return su
}

// AddAge adds i to the "age" field.
func (su *StudentUpdate) AddAge(i int16) *StudentUpdate {
	su.mutation.AddAge(i)
	return su
}

// SetAddress sets the "address" field.
func (su *StudentUpdate) SetAddress(s string) *StudentUpdate {
	su.mutation.SetAddress(s)
	return su
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (su *StudentUpdate) SetNillableAddress(s *string) *StudentUpdate {
	if s != nil {
		su.SetAddress(*s)
	}
	return su
}

// ClearAddress clears the value of the "address" field.
func (su *StudentUpdate) ClearAddress() *StudentUpdate {
	su.mutation.ClearAddress()
	return su
}

// SetScore sets the "score" field.
func (su *StudentUpdate) SetScore(i int32) *StudentUpdate {
	su.mutation.ResetScore()
	su.mutation.SetScore(i)
	return su
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (su *StudentUpdate) SetNillableScore(i *int32) *StudentUpdate {
	if i != nil {
		su.SetScore(*i)
	}
	return su
}

// AddScore adds i to the "score" field.
func (su *StudentUpdate) AddScore(i int32) *StudentUpdate {
	su.mutation.AddScore(i)
	return su
}

// ClearScore clears the value of the "score" field.
func (su *StudentUpdate) ClearScore() *StudentUpdate {
	su.mutation.ClearScore()
	return su
}

// SetWeight sets the "weight" field.
func (su *StudentUpdate) SetWeight(u uint32) *StudentUpdate {
	su.mutation.ResetWeight()
	su.mutation.SetWeight(u)
	return su
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (su *StudentUpdate) SetNillableWeight(u *uint32) *StudentUpdate {
	if u != nil {
		su.SetWeight(*u)
	}
	return su
}

// AddWeight adds u to the "weight" field.
func (su *StudentUpdate) AddWeight(u int32) *StudentUpdate {
	su.mutation.AddWeight(u)
	return su
}

// ClearWeight clears the value of the "weight" field.
func (su *StudentUpdate) ClearWeight() *StudentUpdate {
	su.mutation.ClearWeight()
	return su
}

// SetHealthy sets the "healthy" field.
func (su *StudentUpdate) SetHealthy(b bool) *StudentUpdate {
	su.mutation.SetHealthy(b)
	return su
}

// SetNillableHealthy sets the "healthy" field if the given value is not nil.
func (su *StudentUpdate) SetNillableHealthy(b *bool) *StudentUpdate {
	if b != nil {
		su.SetHealthy(*b)
	}
	return su
}

// ClearHealthy clears the value of the "healthy" field.
func (su *StudentUpdate) ClearHealthy() *StudentUpdate {
	su.mutation.ClearHealthy()
	return su
}

// SetCode sets the "code" field.
func (su *StudentUpdate) SetCode(i int64) *StudentUpdate {
	su.mutation.ResetCode()
	su.mutation.SetCode(i)
	return su
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (su *StudentUpdate) SetNillableCode(i *int64) *StudentUpdate {
	if i != nil {
		su.SetCode(*i)
	}
	return su
}

// AddCode adds i to the "code" field.
func (su *StudentUpdate) AddCode(i int64) *StudentUpdate {
	su.mutation.AddCode(i)
	return su
}

// ClearCode clears the value of the "code" field.
func (su *StudentUpdate) ClearCode() *StudentUpdate {
	su.mutation.ClearCode()
	return su
}

// SetIdentifyID sets the "identify_id" field.
func (su *StudentUpdate) SetIdentifyID(s string) *StudentUpdate {
	su.mutation.SetIdentifyID(s)
	return su
}

// SetNillableIdentifyID sets the "identify_id" field if the given value is not nil.
func (su *StudentUpdate) SetNillableIdentifyID(s *string) *StudentUpdate {
	if s != nil {
		su.SetIdentifyID(*s)
	}
	return su
}

// ClearIdentifyID clears the value of the "identify_id" field.
func (su *StudentUpdate) ClearIdentifyID() *StudentUpdate {
	su.mutation.ClearIdentifyID()
	return su
}

// SetHeight sets the "height" field.
func (su *StudentUpdate) SetHeight(i int) *StudentUpdate {
	su.mutation.ResetHeight()
	su.mutation.SetHeight(i)
	return su
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (su *StudentUpdate) SetNillableHeight(i *int) *StudentUpdate {
	if i != nil {
		su.SetHeight(*i)
	}
	return su
}

// AddHeight adds i to the "height" field.
func (su *StudentUpdate) AddHeight(i int) *StudentUpdate {
	su.mutation.AddHeight(i)
	return su
}

// ClearHeight clears the value of the "height" field.
func (su *StudentUpdate) ClearHeight() *StudentUpdate {
	su.mutation.ClearHeight()
	return su
}

// SetExpiredAt sets the "expired_at" field.
func (su *StudentUpdate) SetExpiredAt(t time.Time) *StudentUpdate {
	su.mutation.SetExpiredAt(t)
	return su
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (su *StudentUpdate) SetNillableExpiredAt(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetExpiredAt(*t)
	}
	return su
}

// ClearExpiredAt clears the value of the "expired_at" field.
func (su *StudentUpdate) ClearExpiredAt() *StudentUpdate {
	su.mutation.ClearExpiredAt()
	return su
}

// SetStudentNumber sets the "student_number" field.
func (su *StudentUpdate) SetStudentNumber(u uuid.UUID) *StudentUpdate {
	su.mutation.SetStudentNumber(u)
	return su
}

// SetNillableStudentNumber sets the "student_number" field if the given value is not nil.
func (su *StudentUpdate) SetNillableStudentNumber(u *uuid.UUID) *StudentUpdate {
	if u != nil {
		su.SetStudentNumber(*u)
	}
	return su
}

// ClearStudentNumber clears the value of the "student_number" field.
func (su *StudentUpdate) ClearStudentNumber() *StudentUpdate {
	su.mutation.ClearStudentNumber()
	return su
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (su *StudentUpdate) AddTeacherIDs(ids ...uint64) *StudentUpdate {
	su.mutation.AddTeacherIDs(ids...)
	return su
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (su *StudentUpdate) AddTeachers(t ...*Teacher) *StudentUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTeacherIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (su *StudentUpdate) ClearTeachers() *StudentUpdate {
	su.mutation.ClearTeachers()
	return su
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (su *StudentUpdate) RemoveTeacherIDs(ids ...uint64) *StudentUpdate {
	su.mutation.RemoveTeacherIDs(ids...)
	return su
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (su *StudentUpdate) RemoveTeachers(t ...*Teacher) *StudentUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTeacherIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StudentUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(student.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedStatus(); ok {
		_spec.AddField(student.FieldStatus, field.TypeUint8, value)
	}
	if su.mutation.StatusCleared() {
		_spec.ClearField(student.FieldStatus, field.TypeUint8)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt16, value)
	}
	if value, ok := su.mutation.AddedAge(); ok {
		_spec.AddField(student.FieldAge, field.TypeInt16, value)
	}
	if value, ok := su.mutation.Address(); ok {
		_spec.SetField(student.FieldAddress, field.TypeString, value)
	}
	if su.mutation.AddressCleared() {
		_spec.ClearField(student.FieldAddress, field.TypeString)
	}
	if value, ok := su.mutation.Score(); ok {
		_spec.SetField(student.FieldScore, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedScore(); ok {
		_spec.AddField(student.FieldScore, field.TypeInt32, value)
	}
	if su.mutation.ScoreCleared() {
		_spec.ClearField(student.FieldScore, field.TypeInt32)
	}
	if value, ok := su.mutation.Weight(); ok {
		_spec.SetField(student.FieldWeight, field.TypeUint32, value)
	}
	if value, ok := su.mutation.AddedWeight(); ok {
		_spec.AddField(student.FieldWeight, field.TypeUint32, value)
	}
	if su.mutation.WeightCleared() {
		_spec.ClearField(student.FieldWeight, field.TypeUint32)
	}
	if value, ok := su.mutation.Healthy(); ok {
		_spec.SetField(student.FieldHealthy, field.TypeBool, value)
	}
	if su.mutation.HealthyCleared() {
		_spec.ClearField(student.FieldHealthy, field.TypeBool)
	}
	if value, ok := su.mutation.Code(); ok {
		_spec.SetField(student.FieldCode, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedCode(); ok {
		_spec.AddField(student.FieldCode, field.TypeInt64, value)
	}
	if su.mutation.CodeCleared() {
		_spec.ClearField(student.FieldCode, field.TypeInt64)
	}
	if value, ok := su.mutation.IdentifyID(); ok {
		_spec.SetField(student.FieldIdentifyID, field.TypeString, value)
	}
	if su.mutation.IdentifyIDCleared() {
		_spec.ClearField(student.FieldIdentifyID, field.TypeString)
	}
	if value, ok := su.mutation.Height(); ok {
		_spec.SetField(student.FieldHeight, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedHeight(); ok {
		_spec.AddField(student.FieldHeight, field.TypeInt, value)
	}
	if su.mutation.HeightCleared() {
		_spec.ClearField(student.FieldHeight, field.TypeInt)
	}
	if value, ok := su.mutation.ExpiredAt(); ok {
		_spec.SetField(student.FieldExpiredAt, field.TypeTime, value)
	}
	if su.mutation.ExpiredAtCleared() {
		_spec.ClearField(student.FieldExpiredAt, field.TypeTime)
	}
	if value, ok := su.mutation.StudentNumber(); ok {
		_spec.SetField(student.FieldStudentNumber, field.TypeUUID, value)
	}
	if su.mutation.StudentNumberCleared() {
		_spec.ClearField(student.FieldStudentNumber, field.TypeUUID)
	}
	if su.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !su.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StudentUpdateOne) SetUpdatedAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetStatus sets the "status" field.
func (suo *StudentUpdateOne) SetStatus(u uint8) *StudentUpdateOne {
	suo.mutation.ResetStatus()
	suo.mutation.SetStatus(u)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableStatus(u *uint8) *StudentUpdateOne {
	if u != nil {
		suo.SetStatus(*u)
	}
	return suo
}

// AddStatus adds u to the "status" field.
func (suo *StudentUpdateOne) AddStatus(u int8) *StudentUpdateOne {
	suo.mutation.AddStatus(u)
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *StudentUpdateOne) ClearStatus() *StudentUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableName(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetAge sets the "age" field.
func (suo *StudentUpdateOne) SetAge(i int16) *StudentUpdateOne {
	suo.mutation.ResetAge()
	suo.mutation.SetAge(i)
	return suo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableAge(i *int16) *StudentUpdateOne {
	if i != nil {
		suo.SetAge(*i)
	}
	return suo
}

// AddAge adds i to the "age" field.
func (suo *StudentUpdateOne) AddAge(i int16) *StudentUpdateOne {
	suo.mutation.AddAge(i)
	return suo
}

// SetAddress sets the "address" field.
func (suo *StudentUpdateOne) SetAddress(s string) *StudentUpdateOne {
	suo.mutation.SetAddress(s)
	return suo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableAddress(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetAddress(*s)
	}
	return suo
}

// ClearAddress clears the value of the "address" field.
func (suo *StudentUpdateOne) ClearAddress() *StudentUpdateOne {
	suo.mutation.ClearAddress()
	return suo
}

// SetScore sets the "score" field.
func (suo *StudentUpdateOne) SetScore(i int32) *StudentUpdateOne {
	suo.mutation.ResetScore()
	suo.mutation.SetScore(i)
	return suo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableScore(i *int32) *StudentUpdateOne {
	if i != nil {
		suo.SetScore(*i)
	}
	return suo
}

// AddScore adds i to the "score" field.
func (suo *StudentUpdateOne) AddScore(i int32) *StudentUpdateOne {
	suo.mutation.AddScore(i)
	return suo
}

// ClearScore clears the value of the "score" field.
func (suo *StudentUpdateOne) ClearScore() *StudentUpdateOne {
	suo.mutation.ClearScore()
	return suo
}

// SetWeight sets the "weight" field.
func (suo *StudentUpdateOne) SetWeight(u uint32) *StudentUpdateOne {
	suo.mutation.ResetWeight()
	suo.mutation.SetWeight(u)
	return suo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableWeight(u *uint32) *StudentUpdateOne {
	if u != nil {
		suo.SetWeight(*u)
	}
	return suo
}

// AddWeight adds u to the "weight" field.
func (suo *StudentUpdateOne) AddWeight(u int32) *StudentUpdateOne {
	suo.mutation.AddWeight(u)
	return suo
}

// ClearWeight clears the value of the "weight" field.
func (suo *StudentUpdateOne) ClearWeight() *StudentUpdateOne {
	suo.mutation.ClearWeight()
	return suo
}

// SetHealthy sets the "healthy" field.
func (suo *StudentUpdateOne) SetHealthy(b bool) *StudentUpdateOne {
	suo.mutation.SetHealthy(b)
	return suo
}

// SetNillableHealthy sets the "healthy" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableHealthy(b *bool) *StudentUpdateOne {
	if b != nil {
		suo.SetHealthy(*b)
	}
	return suo
}

// ClearHealthy clears the value of the "healthy" field.
func (suo *StudentUpdateOne) ClearHealthy() *StudentUpdateOne {
	suo.mutation.ClearHealthy()
	return suo
}

// SetCode sets the "code" field.
func (suo *StudentUpdateOne) SetCode(i int64) *StudentUpdateOne {
	suo.mutation.ResetCode()
	suo.mutation.SetCode(i)
	return suo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableCode(i *int64) *StudentUpdateOne {
	if i != nil {
		suo.SetCode(*i)
	}
	return suo
}

// AddCode adds i to the "code" field.
func (suo *StudentUpdateOne) AddCode(i int64) *StudentUpdateOne {
	suo.mutation.AddCode(i)
	return suo
}

// ClearCode clears the value of the "code" field.
func (suo *StudentUpdateOne) ClearCode() *StudentUpdateOne {
	suo.mutation.ClearCode()
	return suo
}

// SetIdentifyID sets the "identify_id" field.
func (suo *StudentUpdateOne) SetIdentifyID(s string) *StudentUpdateOne {
	suo.mutation.SetIdentifyID(s)
	return suo
}

// SetNillableIdentifyID sets the "identify_id" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableIdentifyID(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetIdentifyID(*s)
	}
	return suo
}

// ClearIdentifyID clears the value of the "identify_id" field.
func (suo *StudentUpdateOne) ClearIdentifyID() *StudentUpdateOne {
	suo.mutation.ClearIdentifyID()
	return suo
}

// SetHeight sets the "height" field.
func (suo *StudentUpdateOne) SetHeight(i int) *StudentUpdateOne {
	suo.mutation.ResetHeight()
	suo.mutation.SetHeight(i)
	return suo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableHeight(i *int) *StudentUpdateOne {
	if i != nil {
		suo.SetHeight(*i)
	}
	return suo
}

// AddHeight adds i to the "height" field.
func (suo *StudentUpdateOne) AddHeight(i int) *StudentUpdateOne {
	suo.mutation.AddHeight(i)
	return suo
}

// ClearHeight clears the value of the "height" field.
func (suo *StudentUpdateOne) ClearHeight() *StudentUpdateOne {
	suo.mutation.ClearHeight()
	return suo
}

// SetExpiredAt sets the "expired_at" field.
func (suo *StudentUpdateOne) SetExpiredAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetExpiredAt(t)
	return suo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableExpiredAt(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetExpiredAt(*t)
	}
	return suo
}

// ClearExpiredAt clears the value of the "expired_at" field.
func (suo *StudentUpdateOne) ClearExpiredAt() *StudentUpdateOne {
	suo.mutation.ClearExpiredAt()
	return suo
}

// SetStudentNumber sets the "student_number" field.
func (suo *StudentUpdateOne) SetStudentNumber(u uuid.UUID) *StudentUpdateOne {
	suo.mutation.SetStudentNumber(u)
	return suo
}

// SetNillableStudentNumber sets the "student_number" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableStudentNumber(u *uuid.UUID) *StudentUpdateOne {
	if u != nil {
		suo.SetStudentNumber(*u)
	}
	return suo
}

// ClearStudentNumber clears the value of the "student_number" field.
func (suo *StudentUpdateOne) ClearStudentNumber() *StudentUpdateOne {
	suo.mutation.ClearStudentNumber()
	return suo
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (suo *StudentUpdateOne) AddTeacherIDs(ids ...uint64) *StudentUpdateOne {
	suo.mutation.AddTeacherIDs(ids...)
	return suo
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (suo *StudentUpdateOne) AddTeachers(t ...*Teacher) *StudentUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTeacherIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (suo *StudentUpdateOne) ClearTeachers() *StudentUpdateOne {
	suo.mutation.ClearTeachers()
	return suo
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (suo *StudentUpdateOne) RemoveTeacherIDs(ids ...uint64) *StudentUpdateOne {
	suo.mutation.RemoveTeacherIDs(ids...)
	return suo
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (suo *StudentUpdateOne) RemoveTeachers(t ...*Teacher) *StudentUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTeacherIDs(ids...)
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StudentUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(student.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedStatus(); ok {
		_spec.AddField(student.FieldStatus, field.TypeUint8, value)
	}
	if suo.mutation.StatusCleared() {
		_spec.ClearField(student.FieldStatus, field.TypeUint8)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt16, value)
	}
	if value, ok := suo.mutation.AddedAge(); ok {
		_spec.AddField(student.FieldAge, field.TypeInt16, value)
	}
	if value, ok := suo.mutation.Address(); ok {
		_spec.SetField(student.FieldAddress, field.TypeString, value)
	}
	if suo.mutation.AddressCleared() {
		_spec.ClearField(student.FieldAddress, field.TypeString)
	}
	if value, ok := suo.mutation.Score(); ok {
		_spec.SetField(student.FieldScore, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedScore(); ok {
		_spec.AddField(student.FieldScore, field.TypeInt32, value)
	}
	if suo.mutation.ScoreCleared() {
		_spec.ClearField(student.FieldScore, field.TypeInt32)
	}
	if value, ok := suo.mutation.Weight(); ok {
		_spec.SetField(student.FieldWeight, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.AddedWeight(); ok {
		_spec.AddField(student.FieldWeight, field.TypeUint32, value)
	}
	if suo.mutation.WeightCleared() {
		_spec.ClearField(student.FieldWeight, field.TypeUint32)
	}
	if value, ok := suo.mutation.Healthy(); ok {
		_spec.SetField(student.FieldHealthy, field.TypeBool, value)
	}
	if suo.mutation.HealthyCleared() {
		_spec.ClearField(student.FieldHealthy, field.TypeBool)
	}
	if value, ok := suo.mutation.Code(); ok {
		_spec.SetField(student.FieldCode, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedCode(); ok {
		_spec.AddField(student.FieldCode, field.TypeInt64, value)
	}
	if suo.mutation.CodeCleared() {
		_spec.ClearField(student.FieldCode, field.TypeInt64)
	}
	if value, ok := suo.mutation.IdentifyID(); ok {
		_spec.SetField(student.FieldIdentifyID, field.TypeString, value)
	}
	if suo.mutation.IdentifyIDCleared() {
		_spec.ClearField(student.FieldIdentifyID, field.TypeString)
	}
	if value, ok := suo.mutation.Height(); ok {
		_spec.SetField(student.FieldHeight, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedHeight(); ok {
		_spec.AddField(student.FieldHeight, field.TypeInt, value)
	}
	if suo.mutation.HeightCleared() {
		_spec.ClearField(student.FieldHeight, field.TypeInt)
	}
	if value, ok := suo.mutation.ExpiredAt(); ok {
		_spec.SetField(student.FieldExpiredAt, field.TypeTime, value)
	}
	if suo.mutation.ExpiredAtCleared() {
		_spec.ClearField(student.FieldExpiredAt, field.TypeTime)
	}
	if value, ok := suo.mutation.StudentNumber(); ok {
		_spec.SetField(student.FieldStudentNumber, field.TypeUUID, value)
	}
	if suo.mutation.StudentNumberCleared() {
		_spec.ClearField(student.FieldStudentNumber, field.TypeUUID)
	}
	if suo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !suo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   student.TeachersTable,
			Columns: student.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
