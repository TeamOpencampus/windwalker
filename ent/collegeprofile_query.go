// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"windwalker/ent/collegeprofile"
	"windwalker/ent/predicate"
	"windwalker/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// CollegeProfileQuery is the builder for querying CollegeProfile entities.
type CollegeProfileQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CollegeProfile
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CollegeProfileQuery builder.
func (cpq *CollegeProfileQuery) Where(ps ...predicate.CollegeProfile) *CollegeProfileQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *CollegeProfileQuery) Limit(limit int) *CollegeProfileQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *CollegeProfileQuery) Offset(offset int) *CollegeProfileQuery {
	cpq.offset = &offset
	return cpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpq *CollegeProfileQuery) Unique(unique bool) *CollegeProfileQuery {
	cpq.unique = &unique
	return cpq
}

// Order adds an order step to the query.
func (cpq *CollegeProfileQuery) Order(o ...OrderFunc) *CollegeProfileQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// QueryUser chains the current query on the "user" edge.
func (cpq *CollegeProfileQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(collegeprofile.Table, collegeprofile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, collegeprofile.UserTable, collegeprofile.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CollegeProfile entity from the query.
// Returns a *NotFoundError when no CollegeProfile was found.
func (cpq *CollegeProfileQuery) First(ctx context.Context) (*CollegeProfile, error) {
	nodes, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{collegeprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *CollegeProfileQuery) FirstX(ctx context.Context) *CollegeProfile {
	node, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CollegeProfile ID from the query.
// Returns a *NotFoundError when no CollegeProfile ID was found.
func (cpq *CollegeProfileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{collegeprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpq *CollegeProfileQuery) FirstIDX(ctx context.Context) string {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CollegeProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CollegeProfile entity is found.
// Returns a *NotFoundError when no CollegeProfile entities are found.
func (cpq *CollegeProfileQuery) Only(ctx context.Context) (*CollegeProfile, error) {
	nodes, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{collegeprofile.Label}
	default:
		return nil, &NotSingularError{collegeprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *CollegeProfileQuery) OnlyX(ctx context.Context) *CollegeProfile {
	node, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CollegeProfile ID in the query.
// Returns a *NotSingularError when more than one CollegeProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (cpq *CollegeProfileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = &NotSingularError{collegeprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpq *CollegeProfileQuery) OnlyIDX(ctx context.Context) string {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CollegeProfiles.
func (cpq *CollegeProfileQuery) All(ctx context.Context) ([]*CollegeProfile, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *CollegeProfileQuery) AllX(ctx context.Context) []*CollegeProfile {
	nodes, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CollegeProfile IDs.
func (cpq *CollegeProfileQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := cpq.Select(collegeprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *CollegeProfileQuery) IDsX(ctx context.Context) []string {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *CollegeProfileQuery) Count(ctx context.Context) (int, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *CollegeProfileQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *CollegeProfileQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *CollegeProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CollegeProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *CollegeProfileQuery) Clone() *CollegeProfileQuery {
	if cpq == nil {
		return nil
	}
	return &CollegeProfileQuery{
		config:     cpq.config,
		limit:      cpq.limit,
		offset:     cpq.offset,
		order:      append([]OrderFunc{}, cpq.order...),
		predicates: append([]predicate.CollegeProfile{}, cpq.predicates...),
		withUser:   cpq.withUser.Clone(),
		// clone intermediate query.
		sql:    cpq.sql.Clone(),
		path:   cpq.path,
		unique: cpq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CollegeProfileQuery) WithUser(opts ...func(*UserQuery)) *CollegeProfileQuery {
	query := &UserQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withUser = query
	return cpq
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
//	client.CollegeProfile.Query().
//		GroupBy(collegeprofile.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cpq *CollegeProfileQuery) GroupBy(field string, fields ...string) *CollegeProfileGroupBy {
	group := &CollegeProfileGroupBy{config: cpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpq.sqlQuery(ctx), nil
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
//	client.CollegeProfile.Query().
//		Select(collegeprofile.FieldName).
//		Scan(ctx, &v)
//
func (cpq *CollegeProfileQuery) Select(fields ...string) *CollegeProfileSelect {
	cpq.fields = append(cpq.fields, fields...)
	return &CollegeProfileSelect{CollegeProfileQuery: cpq}
}

func (cpq *CollegeProfileQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpq.fields {
		if !collegeprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cpq.path != nil {
		prev, err := cpq.path(ctx)
		if err != nil {
			return err
		}
		cpq.sql = prev
	}
	return nil
}

func (cpq *CollegeProfileQuery) sqlAll(ctx context.Context) ([]*CollegeProfile, error) {
	var (
		nodes       = []*CollegeProfile{}
		withFKs     = cpq.withFKs
		_spec       = cpq.querySpec()
		loadedTypes = [1]bool{
			cpq.withUser != nil,
		}
	)
	if cpq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, collegeprofile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CollegeProfile{config: cpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, cpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cpq.withUser; query != nil {
		ids := make([]xid.ID, 0, len(nodes))
		nodeids := make(map[xid.ID][]*CollegeProfile)
		for i := range nodes {
			if nodes[i].user_college_profile == nil {
				continue
			}
			fk := *nodes[i].user_college_profile
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_college_profile" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (cpq *CollegeProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpq.querySpec()
	_spec.Node.Columns = cpq.fields
	if len(cpq.fields) > 0 {
		_spec.Unique = cpq.unique != nil && *cpq.unique
	}
	return sqlgraph.CountNodes(ctx, cpq.driver, _spec)
}

func (cpq *CollegeProfileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpq *CollegeProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collegeprofile.Table,
			Columns: collegeprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: collegeprofile.FieldID,
			},
		},
		From:   cpq.sql,
		Unique: true,
	}
	if unique := cpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collegeprofile.FieldID)
		for i := range fields {
			if fields[i] != collegeprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpq *CollegeProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(collegeprofile.Table)
	columns := cpq.fields
	if len(columns) == 0 {
		columns = collegeprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cpq.unique != nil && *cpq.unique {
		selector.Distinct()
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CollegeProfileGroupBy is the group-by builder for CollegeProfile entities.
type CollegeProfileGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *CollegeProfileGroupBy) Aggregate(fns ...AggregateFunc) *CollegeProfileGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpgb *CollegeProfileGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpgb.path(ctx)
	if err != nil {
		return err
	}
	cpgb.sql = query
	return cpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) StringsX(ctx context.Context) []string {
	v, err := cpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) StringX(ctx context.Context) string {
	v, err := cpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) IntsX(ctx context.Context) []int {
	v, err := cpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) IntX(ctx context.Context) int {
	v, err := cpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CollegeProfileGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cpgb *CollegeProfileGroupBy) BoolX(ctx context.Context) bool {
	v, err := cpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpgb *CollegeProfileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpgb.fields {
		if !collegeprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *CollegeProfileGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql.Select()
	aggregation := make([]string, 0, len(cpgb.fns))
	for _, fn := range cpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
		for _, f := range cpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpgb.fields...)...)
}

// CollegeProfileSelect is the builder for selecting fields of CollegeProfile entities.
type CollegeProfileSelect struct {
	*CollegeProfileQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cps *CollegeProfileSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cps.prepareQuery(ctx); err != nil {
		return err
	}
	cps.sql = cps.CollegeProfileQuery.sqlQuery(ctx)
	return cps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cps *CollegeProfileSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cps *CollegeProfileSelect) StringsX(ctx context.Context) []string {
	v, err := cps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cps *CollegeProfileSelect) StringX(ctx context.Context) string {
	v, err := cps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cps *CollegeProfileSelect) IntsX(ctx context.Context) []int {
	v, err := cps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cps *CollegeProfileSelect) IntX(ctx context.Context) int {
	v, err := cps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cps *CollegeProfileSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cps *CollegeProfileSelect) Float64X(ctx context.Context) float64 {
	v, err := cps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CollegeProfileSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cps *CollegeProfileSelect) BoolsX(ctx context.Context) []bool {
	v, err := cps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cps *CollegeProfileSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{collegeprofile.Label}
	default:
		err = fmt.Errorf("ent: CollegeProfileSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cps *CollegeProfileSelect) BoolX(ctx context.Context) bool {
	v, err := cps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cps *CollegeProfileSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sql.Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
