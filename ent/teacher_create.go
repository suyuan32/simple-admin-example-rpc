// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
)

// TeacherCreate is the builder for creating a Teacher entity.
type TeacherCreate struct {
	config
	mutation *TeacherMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TeacherCreate) SetCreatedAt(t time.Time) *TeacherCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TeacherCreate) SetNillableCreatedAt(t *time.Time) *TeacherCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TeacherCreate) SetUpdatedAt(t time.Time) *TeacherCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TeacherCreate) SetNillableUpdatedAt(t *time.Time) *TeacherCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TeacherCreate) SetName(s string) *TeacherCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetAge sets the "age" field.
func (tc *TeacherCreate) SetAge(i int16) *TeacherCreate {
	tc.mutation.SetAge(i)
	return tc
}

// SetID sets the "id" field.
func (tc *TeacherCreate) SetID(u uint64) *TeacherCreate {
	tc.mutation.SetID(u)
	return tc
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (tc *TeacherCreate) AddStudentIDs(ids ...uuid.UUID) *TeacherCreate {
	tc.mutation.AddStudentIDs(ids...)
	return tc
}

// AddStudents adds the "students" edges to the Student entity.
func (tc *TeacherCreate) AddStudents(s ...*Student) *TeacherCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddStudentIDs(ids...)
}

// Mutation returns the TeacherMutation object of the builder.
func (tc *TeacherCreate) Mutation() *TeacherMutation {
	return tc.mutation
}

// Save creates the Teacher in the database.
func (tc *TeacherCreate) Save(ctx context.Context) (*Teacher, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeacherCreate) SaveX(ctx context.Context) *Teacher {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeacherCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeacherCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TeacherCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := teacher.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := teacher.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeacherCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Teacher.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Teacher.updated_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Teacher.name"`)}
	}
	if _, ok := tc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Teacher.age"`)}
	}
	if len(tc.mutation.StudentsIDs()) == 0 {
		return &ValidationError{Name: "students", err: errors.New(`ent: missing required edge "Teacher.students"`)}
	}
	return nil
}

func (tc *TeacherCreate) sqlSave(ctx context.Context) (*Teacher, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TeacherCreate) createSpec() (*Teacher, *sqlgraph.CreateSpec) {
	var (
		_node = &Teacher{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(teacher.Table, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(teacher.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(teacher.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Age(); ok {
		_spec.SetField(teacher.FieldAge, field.TypeInt16, value)
		_node.Age = value
	}
	if nodes := tc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   teacher.StudentsTable,
			Columns: teacher.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TeacherCreateBulk is the builder for creating many Teacher entities in bulk.
type TeacherCreateBulk struct {
	config
	err      error
	builders []*TeacherCreate
}

// Save creates the Teacher entities in the database.
func (tcb *TeacherCreateBulk) Save(ctx context.Context) ([]*Teacher, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Teacher, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeacherMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeacherCreateBulk) SaveX(ctx context.Context) []*Teacher {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeacherCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeacherCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
