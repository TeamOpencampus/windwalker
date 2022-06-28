// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"windwalker/ent/predicate"
	"windwalker/ent/studentacademicprofile"
	"windwalker/ent/studentprofile"
	"windwalker/ent/studentworkprofile"
	"windwalker/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// StudentProfileQuery is the builder for querying StudentProfile entities.
type StudentProfileQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StudentProfile
	// eager-loading edges.
	withUser            *UserQuery
	withAcademicProfile *StudentAcademicProfileQuery
	withWorkProfile     *StudentWorkProfileQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StudentProfileQuery builder.
func (spq *StudentProfileQuery) Where(ps ...predicate.StudentProfile) *StudentProfileQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit adds a limit step to the query.
func (spq *StudentProfileQuery) Limit(limit int) *StudentProfileQuery {
	spq.limit = &limit
	return spq
}

// Offset adds an offset step to the query.
func (spq *StudentProfileQuery) Offset(offset int) *StudentProfileQuery {
	spq.offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *StudentProfileQuery) Unique(unique bool) *StudentProfileQuery {
	spq.unique = &unique
	return spq
}

// Order adds an order step to the query.
func (spq *StudentProfileQuery) Order(o ...OrderFunc) *StudentProfileQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryUser chains the current query on the "user" edge.
func (spq *StudentProfileQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: spq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studentprofile.Table, studentprofile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, studentprofile.UserTable, studentprofile.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAcademicProfile chains the current query on the "academic_profile" edge.
func (spq *StudentProfileQuery) QueryAcademicProfile() *StudentAcademicProfileQuery {
	query := &StudentAcademicProfileQuery{config: spq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studentprofile.Table, studentprofile.FieldID, selector),
			sqlgraph.To(studentacademicprofile.Table, studentacademicprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, studentprofile.AcademicProfileTable, studentprofile.AcademicProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkProfile chains the current query on the "work_profile" edge.
func (spq *StudentProfileQuery) QueryWorkProfile() *StudentWorkProfileQuery {
	query := &StudentWorkProfileQuery{config: spq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studentprofile.Table, studentprofile.FieldID, selector),
			sqlgraph.To(studentworkprofile.Table, studentworkprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, studentprofile.WorkProfileTable, studentprofile.WorkProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StudentProfile entity from the query.
// Returns a *NotFoundError when no StudentProfile was found.
func (spq *StudentProfileQuery) First(ctx context.Context) (*StudentProfile, error) {
	nodes, err := spq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{studentprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *StudentProfileQuery) FirstX(ctx context.Context) *StudentProfile {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StudentProfile ID from the query.
// Returns a *NotFoundError when no StudentProfile ID was found.
func (spq *StudentProfileQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = spq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{studentprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *StudentProfileQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StudentProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StudentProfile entity is found.
// Returns a *NotFoundError when no StudentProfile entities are found.
func (spq *StudentProfileQuery) Only(ctx context.Context) (*StudentProfile, error) {
	nodes, err := spq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{studentprofile.Label}
	default:
		return nil, &NotSingularError{studentprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *StudentProfileQuery) OnlyX(ctx context.Context) *StudentProfile {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StudentProfile ID in the query.
// Returns a *NotSingularError when more than one StudentProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *StudentProfileQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = spq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = &NotSingularError{studentprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *StudentProfileQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StudentProfiles.
func (spq *StudentProfileQuery) All(ctx context.Context) ([]*StudentProfile, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return spq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (spq *StudentProfileQuery) AllX(ctx context.Context) []*StudentProfile {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StudentProfile IDs.
func (spq *StudentProfileQuery) IDs(ctx context.Context) ([]xid.ID, error) {
	var ids []xid.ID
	if err := spq.Select(studentprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *StudentProfileQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *StudentProfileQuery) Count(ctx context.Context) (int, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return spq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (spq *StudentProfileQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *StudentProfileQuery) Exist(ctx context.Context) (bool, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return spq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *StudentProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StudentProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *StudentProfileQuery) Clone() *StudentProfileQuery {
	if spq == nil {
		return nil
	}
	return &StudentProfileQuery{
		config:              spq.config,
		limit:               spq.limit,
		offset:              spq.offset,
		order:               append([]OrderFunc{}, spq.order...),
		predicates:          append([]predicate.StudentProfile{}, spq.predicates...),
		withUser:            spq.withUser.Clone(),
		withAcademicProfile: spq.withAcademicProfile.Clone(),
		withWorkProfile:     spq.withWorkProfile.Clone(),
		// clone intermediate query.
		sql:    spq.sql.Clone(),
		path:   spq.path,
		unique: spq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StudentProfileQuery) WithUser(opts ...func(*UserQuery)) *StudentProfileQuery {
	query := &UserQuery{config: spq.config}
	for _, opt := range opts {
		opt(query)
	}
	spq.withUser = query
	return spq
}

// WithAcademicProfile tells the query-builder to eager-load the nodes that are connected to
// the "academic_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StudentProfileQuery) WithAcademicProfile(opts ...func(*StudentAcademicProfileQuery)) *StudentProfileQuery {
	query := &StudentAcademicProfileQuery{config: spq.config}
	for _, opt := range opts {
		opt(query)
	}
	spq.withAcademicProfile = query
	return spq
}

// WithWorkProfile tells the query-builder to eager-load the nodes that are connected to
// the "work_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StudentProfileQuery) WithWorkProfile(opts ...func(*StudentWorkProfileQuery)) *StudentProfileQuery {
	query := &StudentWorkProfileQuery{config: spq.config}
	for _, opt := range opts {
		opt(query)
	}
	spq.withWorkProfile = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StudentProfile.Query().
//		GroupBy(studentprofile.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (spq *StudentProfileQuery) GroupBy(field string, fields ...string) *StudentProfileGroupBy {
	group := &StudentProfileGroupBy{config: spq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return spq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.StudentProfile.Query().
//		Select(studentprofile.FieldName).
//		Scan(ctx, &v)
//
func (spq *StudentProfileQuery) Select(fields ...string) *StudentProfileSelect {
	spq.fields = append(spq.fields, fields...)
	return &StudentProfileSelect{StudentProfileQuery: spq}
}

func (spq *StudentProfileQuery) prepareQuery(ctx context.Context) error {
	for _, f := range spq.fields {
		if !studentprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *StudentProfileQuery) sqlAll(ctx context.Context) ([]*StudentProfile, error) {
	var (
		nodes       = []*StudentProfile{}
		withFKs     = spq.withFKs
		_spec       = spq.querySpec()
		loadedTypes = [3]bool{
			spq.withUser != nil,
			spq.withAcademicProfile != nil,
			spq.withWorkProfile != nil,
		}
	)
	if spq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, studentprofile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StudentProfile{config: spq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := spq.withUser; query != nil {
		ids := make([]xid.ID, 0, len(nodes))
		nodeids := make(map[xid.ID][]*StudentProfile)
		for i := range nodes {
			if nodes[i].user_student_profile == nil {
				continue
			}
			fk := *nodes[i].user_student_profile
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_student_profile" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := spq.withAcademicProfile; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[xid.ID]*StudentProfile)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.AcademicProfile = []*StudentAcademicProfile{}
		}
		query.withFKs = true
		query.Where(predicate.StudentAcademicProfile(func(s *sql.Selector) {
			s.Where(sql.InValues(studentprofile.AcademicProfileColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.student_profile_academic_profile
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "student_profile_academic_profile" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "student_profile_academic_profile" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.AcademicProfile = append(node.Edges.AcademicProfile, n)
		}
	}

	if query := spq.withWorkProfile; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[xid.ID]*StudentProfile)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.WorkProfile = []*StudentWorkProfile{}
		}
		query.withFKs = true
		query.Where(predicate.StudentWorkProfile(func(s *sql.Selector) {
			s.Where(sql.InValues(studentprofile.WorkProfileColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.student_profile_work_profile
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "student_profile_work_profile" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "student_profile_work_profile" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.WorkProfile = append(node.Edges.WorkProfile, n)
		}
	}

	return nodes, nil
}

func (spq *StudentProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.fields
	if len(spq.fields) > 0 {
		_spec.Unique = spq.unique != nil && *spq.unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *StudentProfileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := spq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (spq *StudentProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studentprofile.Table,
			Columns: studentprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentprofile.FieldID,
			},
		},
		From:   spq.sql,
		Unique: true,
	}
	if unique := spq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := spq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studentprofile.FieldID)
		for i := range fields {
			if fields[i] != studentprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *StudentProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(studentprofile.Table)
	columns := spq.fields
	if len(columns) == 0 {
		columns = studentprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.unique != nil && *spq.unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StudentProfileGroupBy is the group-by builder for StudentProfile entities.
type StudentProfileGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *StudentProfileGroupBy) Aggregate(fns ...AggregateFunc) *StudentProfileGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the group-by query and scans the result into the given value.
func (spgb *StudentProfileGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := spgb.path(ctx)
	if err != nil {
		return err
	}
	spgb.sql = query
	return spgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := spgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: StudentProfileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) StringsX(ctx context.Context) []string {
	v, err := spgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = spgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) StringX(ctx context.Context) string {
	v, err := spgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: StudentProfileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) IntsX(ctx context.Context) []int {
	v, err := spgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = spgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) IntX(ctx context.Context) int {
	v, err := spgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: StudentProfileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := spgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = spgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) Float64X(ctx context.Context) float64 {
	v, err := spgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: StudentProfileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := spgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *StudentProfileGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = spgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (spgb *StudentProfileGroupBy) BoolX(ctx context.Context) bool {
	v, err := spgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (spgb *StudentProfileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range spgb.fields {
		if !studentprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := spgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (spgb *StudentProfileGroupBy) sqlQuery() *sql.Selector {
	selector := spgb.sql.Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(spgb.fields)+len(spgb.fns))
		for _, f := range spgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(spgb.fields...)...)
}

// StudentProfileSelect is the builder for selecting fields of StudentProfile entities.
type StudentProfileSelect struct {
	*StudentProfileQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sps *StudentProfileSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	sps.sql = sps.StudentProfileQuery.sqlQuery(ctx)
	return sps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sps *StudentProfileSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: StudentProfileSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sps *StudentProfileSelect) StringsX(ctx context.Context) []string {
	v, err := sps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sps *StudentProfileSelect) StringX(ctx context.Context) string {
	v, err := sps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: StudentProfileSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sps *StudentProfileSelect) IntsX(ctx context.Context) []int {
	v, err := sps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sps *StudentProfileSelect) IntX(ctx context.Context) int {
	v, err := sps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: StudentProfileSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sps *StudentProfileSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sps *StudentProfileSelect) Float64X(ctx context.Context) float64 {
	v, err := sps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: StudentProfileSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sps *StudentProfileSelect) BoolsX(ctx context.Context) []bool {
	v, err := sps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sps *StudentProfileSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studentprofile.Label}
	default:
		err = fmt.Errorf("ent: StudentProfileSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sps *StudentProfileSelect) BoolX(ctx context.Context) bool {
	v, err := sps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sps *StudentProfileSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sps.sql.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}