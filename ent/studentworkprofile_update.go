// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"windwalker/ent/predicate"
	"windwalker/ent/studentprofile"
	"windwalker/ent/studentworkprofile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// StudentWorkProfileUpdate is the builder for updating StudentWorkProfile entities.
type StudentWorkProfileUpdate struct {
	config
	hooks    []Hook
	mutation *StudentWorkProfileMutation
}

// Where appends a list predicates to the StudentWorkProfileUpdate builder.
func (swpu *StudentWorkProfileUpdate) Where(ps ...predicate.StudentWorkProfile) *StudentWorkProfileUpdate {
	swpu.mutation.Where(ps...)
	return swpu
}

// SetStudentProfileID sets the "student_profile" edge to the StudentProfile entity by ID.
func (swpu *StudentWorkProfileUpdate) SetStudentProfileID(id xid.ID) *StudentWorkProfileUpdate {
	swpu.mutation.SetStudentProfileID(id)
	return swpu
}

// SetNillableStudentProfileID sets the "student_profile" edge to the StudentProfile entity by ID if the given value is not nil.
func (swpu *StudentWorkProfileUpdate) SetNillableStudentProfileID(id *xid.ID) *StudentWorkProfileUpdate {
	if id != nil {
		swpu = swpu.SetStudentProfileID(*id)
	}
	return swpu
}

// SetStudentProfile sets the "student_profile" edge to the StudentProfile entity.
func (swpu *StudentWorkProfileUpdate) SetStudentProfile(s *StudentProfile) *StudentWorkProfileUpdate {
	return swpu.SetStudentProfileID(s.ID)
}

// Mutation returns the StudentWorkProfileMutation object of the builder.
func (swpu *StudentWorkProfileUpdate) Mutation() *StudentWorkProfileMutation {
	return swpu.mutation
}

// ClearStudentProfile clears the "student_profile" edge to the StudentProfile entity.
func (swpu *StudentWorkProfileUpdate) ClearStudentProfile() *StudentWorkProfileUpdate {
	swpu.mutation.ClearStudentProfile()
	return swpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (swpu *StudentWorkProfileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(swpu.hooks) == 0 {
		affected, err = swpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentWorkProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			swpu.mutation = mutation
			affected, err = swpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(swpu.hooks) - 1; i >= 0; i-- {
			if swpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = swpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, swpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (swpu *StudentWorkProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := swpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (swpu *StudentWorkProfileUpdate) Exec(ctx context.Context) error {
	_, err := swpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (swpu *StudentWorkProfileUpdate) ExecX(ctx context.Context) {
	if err := swpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (swpu *StudentWorkProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studentworkprofile.Table,
			Columns: studentworkprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentworkprofile.FieldID,
			},
		},
	}
	if ps := swpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if swpu.mutation.StudentProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentworkprofile.StudentProfileTable,
			Columns: []string{studentworkprofile.StudentProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: studentprofile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := swpu.mutation.StudentProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentworkprofile.StudentProfileTable,
			Columns: []string{studentworkprofile.StudentProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: studentprofile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, swpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studentworkprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StudentWorkProfileUpdateOne is the builder for updating a single StudentWorkProfile entity.
type StudentWorkProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentWorkProfileMutation
}

// SetStudentProfileID sets the "student_profile" edge to the StudentProfile entity by ID.
func (swpuo *StudentWorkProfileUpdateOne) SetStudentProfileID(id xid.ID) *StudentWorkProfileUpdateOne {
	swpuo.mutation.SetStudentProfileID(id)
	return swpuo
}

// SetNillableStudentProfileID sets the "student_profile" edge to the StudentProfile entity by ID if the given value is not nil.
func (swpuo *StudentWorkProfileUpdateOne) SetNillableStudentProfileID(id *xid.ID) *StudentWorkProfileUpdateOne {
	if id != nil {
		swpuo = swpuo.SetStudentProfileID(*id)
	}
	return swpuo
}

// SetStudentProfile sets the "student_profile" edge to the StudentProfile entity.
func (swpuo *StudentWorkProfileUpdateOne) SetStudentProfile(s *StudentProfile) *StudentWorkProfileUpdateOne {
	return swpuo.SetStudentProfileID(s.ID)
}

// Mutation returns the StudentWorkProfileMutation object of the builder.
func (swpuo *StudentWorkProfileUpdateOne) Mutation() *StudentWorkProfileMutation {
	return swpuo.mutation
}

// ClearStudentProfile clears the "student_profile" edge to the StudentProfile entity.
func (swpuo *StudentWorkProfileUpdateOne) ClearStudentProfile() *StudentWorkProfileUpdateOne {
	swpuo.mutation.ClearStudentProfile()
	return swpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (swpuo *StudentWorkProfileUpdateOne) Select(field string, fields ...string) *StudentWorkProfileUpdateOne {
	swpuo.fields = append([]string{field}, fields...)
	return swpuo
}

// Save executes the query and returns the updated StudentWorkProfile entity.
func (swpuo *StudentWorkProfileUpdateOne) Save(ctx context.Context) (*StudentWorkProfile, error) {
	var (
		err  error
		node *StudentWorkProfile
	)
	if len(swpuo.hooks) == 0 {
		node, err = swpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentWorkProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			swpuo.mutation = mutation
			node, err = swpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(swpuo.hooks) - 1; i >= 0; i-- {
			if swpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = swpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, swpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (swpuo *StudentWorkProfileUpdateOne) SaveX(ctx context.Context) *StudentWorkProfile {
	node, err := swpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (swpuo *StudentWorkProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := swpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (swpuo *StudentWorkProfileUpdateOne) ExecX(ctx context.Context) {
	if err := swpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (swpuo *StudentWorkProfileUpdateOne) sqlSave(ctx context.Context) (_node *StudentWorkProfile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studentworkprofile.Table,
			Columns: studentworkprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentworkprofile.FieldID,
			},
		},
	}
	id, ok := swpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StudentWorkProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := swpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studentworkprofile.FieldID)
		for _, f := range fields {
			if !studentworkprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != studentworkprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := swpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if swpuo.mutation.StudentProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentworkprofile.StudentProfileTable,
			Columns: []string{studentworkprofile.StudentProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: studentprofile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := swpuo.mutation.StudentProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentworkprofile.StudentProfileTable,
			Columns: []string{studentworkprofile.StudentProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: studentprofile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudentWorkProfile{config: swpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, swpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studentworkprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
