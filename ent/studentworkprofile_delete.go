// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"windwalker/ent/predicate"
	"windwalker/ent/studentworkprofile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StudentWorkProfileDelete is the builder for deleting a StudentWorkProfile entity.
type StudentWorkProfileDelete struct {
	config
	hooks    []Hook
	mutation *StudentWorkProfileMutation
}

// Where appends a list predicates to the StudentWorkProfileDelete builder.
func (swpd *StudentWorkProfileDelete) Where(ps ...predicate.StudentWorkProfile) *StudentWorkProfileDelete {
	swpd.mutation.Where(ps...)
	return swpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (swpd *StudentWorkProfileDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(swpd.hooks) == 0 {
		affected, err = swpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentWorkProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			swpd.mutation = mutation
			affected, err = swpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(swpd.hooks) - 1; i >= 0; i-- {
			if swpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = swpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, swpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (swpd *StudentWorkProfileDelete) ExecX(ctx context.Context) int {
	n, err := swpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (swpd *StudentWorkProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: studentworkprofile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentworkprofile.FieldID,
			},
		},
	}
	if ps := swpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, swpd.driver, _spec)
}

// StudentWorkProfileDeleteOne is the builder for deleting a single StudentWorkProfile entity.
type StudentWorkProfileDeleteOne struct {
	swpd *StudentWorkProfileDelete
}

// Exec executes the deletion query.
func (swpdo *StudentWorkProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := swpdo.swpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{studentworkprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (swpdo *StudentWorkProfileDeleteOne) ExecX(ctx context.Context) {
	swpdo.swpd.ExecX(ctx)
}
