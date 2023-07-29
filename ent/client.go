// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-rpc/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Student is the client for interacting with the Student builders.
	Student *StudentClient
	// Teacher is the client for interacting with the Teacher builders.
	Teacher *TeacherClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Student = NewStudentClient(c.config)
	c.Teacher = NewTeacherClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Student: NewStudentClient(cfg),
		Teacher: NewTeacherClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Student: NewStudentClient(cfg),
		Teacher: NewTeacherClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Student.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Student.Use(hooks...)
	c.Teacher.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Student.Intercept(interceptors...)
	c.Teacher.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *StudentMutation:
		return c.Student.mutate(ctx, m)
	case *TeacherMutation:
		return c.Teacher.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// StudentClient is a client for the Student schema.
type StudentClient struct {
	config
}

// NewStudentClient returns a client for the Student from the given config.
func NewStudentClient(c config) *StudentClient {
	return &StudentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `student.Hooks(f(g(h())))`.
func (c *StudentClient) Use(hooks ...Hook) {
	c.hooks.Student = append(c.hooks.Student, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `student.Intercept(f(g(h())))`.
func (c *StudentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Student = append(c.inters.Student, interceptors...)
}

// Create returns a builder for creating a Student entity.
func (c *StudentClient) Create() *StudentCreate {
	mutation := newStudentMutation(c.config, OpCreate)
	return &StudentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Student entities.
func (c *StudentClient) CreateBulk(builders ...*StudentCreate) *StudentCreateBulk {
	return &StudentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Student.
func (c *StudentClient) Update() *StudentUpdate {
	mutation := newStudentMutation(c.config, OpUpdate)
	return &StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudentClient) UpdateOne(s *Student) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudent(s))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudentClient) UpdateOneID(id uint64) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudentID(id))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Student.
func (c *StudentClient) Delete() *StudentDelete {
	mutation := newStudentMutation(c.config, OpDelete)
	return &StudentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StudentClient) DeleteOne(s *Student) *StudentDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StudentClient) DeleteOneID(id uint64) *StudentDeleteOne {
	builder := c.Delete().Where(student.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudentDeleteOne{builder}
}

// Query returns a query builder for Student.
func (c *StudentClient) Query() *StudentQuery {
	return &StudentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStudent},
		inters: c.Interceptors(),
	}
}

// Get returns a Student entity by its id.
func (c *StudentClient) Get(ctx context.Context, id uint64) (*Student, error) {
	return c.Query().Where(student.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudentClient) GetX(ctx context.Context, id uint64) *Student {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StudentClient) Hooks() []Hook {
	return c.hooks.Student
}

// Interceptors returns the client interceptors.
func (c *StudentClient) Interceptors() []Interceptor {
	return c.inters.Student
}

func (c *StudentClient) mutate(ctx context.Context, m *StudentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StudentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StudentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Student mutation op: %q", m.Op())
	}
}

// TeacherClient is a client for the Teacher schema.
type TeacherClient struct {
	config
}

// NewTeacherClient returns a client for the Teacher from the given config.
func NewTeacherClient(c config) *TeacherClient {
	return &TeacherClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `teacher.Hooks(f(g(h())))`.
func (c *TeacherClient) Use(hooks ...Hook) {
	c.hooks.Teacher = append(c.hooks.Teacher, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `teacher.Intercept(f(g(h())))`.
func (c *TeacherClient) Intercept(interceptors ...Interceptor) {
	c.inters.Teacher = append(c.inters.Teacher, interceptors...)
}

// Create returns a builder for creating a Teacher entity.
func (c *TeacherClient) Create() *TeacherCreate {
	mutation := newTeacherMutation(c.config, OpCreate)
	return &TeacherCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Teacher entities.
func (c *TeacherClient) CreateBulk(builders ...*TeacherCreate) *TeacherCreateBulk {
	return &TeacherCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Teacher.
func (c *TeacherClient) Update() *TeacherUpdate {
	mutation := newTeacherMutation(c.config, OpUpdate)
	return &TeacherUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeacherClient) UpdateOne(t *Teacher) *TeacherUpdateOne {
	mutation := newTeacherMutation(c.config, OpUpdateOne, withTeacher(t))
	return &TeacherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeacherClient) UpdateOneID(id uuid.UUID) *TeacherUpdateOne {
	mutation := newTeacherMutation(c.config, OpUpdateOne, withTeacherID(id))
	return &TeacherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Teacher.
func (c *TeacherClient) Delete() *TeacherDelete {
	mutation := newTeacherMutation(c.config, OpDelete)
	return &TeacherDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TeacherClient) DeleteOne(t *Teacher) *TeacherDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TeacherClient) DeleteOneID(id uuid.UUID) *TeacherDeleteOne {
	builder := c.Delete().Where(teacher.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeacherDeleteOne{builder}
}

// Query returns a query builder for Teacher.
func (c *TeacherClient) Query() *TeacherQuery {
	return &TeacherQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTeacher},
		inters: c.Interceptors(),
	}
}

// Get returns a Teacher entity by its id.
func (c *TeacherClient) Get(ctx context.Context, id uuid.UUID) (*Teacher, error) {
	return c.Query().Where(teacher.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeacherClient) GetX(ctx context.Context, id uuid.UUID) *Teacher {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TeacherClient) Hooks() []Hook {
	return c.hooks.Teacher
}

// Interceptors returns the client interceptors.
func (c *TeacherClient) Interceptors() []Interceptor {
	return c.inters.Teacher
}

func (c *TeacherClient) mutate(ctx context.Context, m *TeacherMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TeacherCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TeacherUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TeacherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TeacherDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Teacher mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Student, Teacher []ent.Hook
	}
	inters struct {
		Student, Teacher []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
