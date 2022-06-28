// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"windwalker/ent/studentacademicprofile"
	"windwalker/ent/studentprofile"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// StudentAcademicProfileCreate is the builder for creating a StudentAcademicProfile entity.
type StudentAcademicProfileCreate struct {
	config
	mutation *StudentAcademicProfileMutation
	hooks    []Hook
}

// SetCourse sets the "course" field.
func (sapc *StudentAcademicProfileCreate) SetCourse(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetCourse(s)
	return sapc
}

// SetInstitute sets the "institute" field.
func (sapc *StudentAcademicProfileCreate) SetInstitute(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetInstitute(s)
	return sapc
}

// SetBoard sets the "board" field.
func (sapc *StudentAcademicProfileCreate) SetBoard(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetBoard(s)
	return sapc
}

// SetRegNo sets the "reg_no" field.
func (sapc *StudentAcademicProfileCreate) SetRegNo(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetRegNo(s)
	return sapc
}

// SetDepartment sets the "department" field.
func (sapc *StudentAcademicProfileCreate) SetDepartment(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetDepartment(s)
	return sapc
}

// SetNillableDepartment sets the "department" field if the given value is not nil.
func (sapc *StudentAcademicProfileCreate) SetNillableDepartment(s *string) *StudentAcademicProfileCreate {
	if s != nil {
		sapc.SetDepartment(*s)
	}
	return sapc
}

// SetStartDate sets the "start_date" field.
func (sapc *StudentAcademicProfileCreate) SetStartDate(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetStartDate(s)
	return sapc
}

// SetEndDate sets the "end_date" field.
func (sapc *StudentAcademicProfileCreate) SetEndDate(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetEndDate(s)
	return sapc
}

// SetMarks sets the "marks" field.
func (sapc *StudentAcademicProfileCreate) SetMarks(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetMarks(s)
	return sapc
}

// SetID sets the "id" field.
func (sapc *StudentAcademicProfileCreate) SetID(s string) *StudentAcademicProfileCreate {
	sapc.mutation.SetID(s)
	return sapc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sapc *StudentAcademicProfileCreate) SetNillableID(s *string) *StudentAcademicProfileCreate {
	if s != nil {
		sapc.SetID(*s)
	}
	return sapc
}

// SetStudentProfileID sets the "student_profile" edge to the StudentProfile entity by ID.
func (sapc *StudentAcademicProfileCreate) SetStudentProfileID(id xid.ID) *StudentAcademicProfileCreate {
	sapc.mutation.SetStudentProfileID(id)
	return sapc
}

// SetNillableStudentProfileID sets the "student_profile" edge to the StudentProfile entity by ID if the given value is not nil.
func (sapc *StudentAcademicProfileCreate) SetNillableStudentProfileID(id *xid.ID) *StudentAcademicProfileCreate {
	if id != nil {
		sapc = sapc.SetStudentProfileID(*id)
	}
	return sapc
}

// SetStudentProfile sets the "student_profile" edge to the StudentProfile entity.
func (sapc *StudentAcademicProfileCreate) SetStudentProfile(s *StudentProfile) *StudentAcademicProfileCreate {
	return sapc.SetStudentProfileID(s.ID)
}

// Mutation returns the StudentAcademicProfileMutation object of the builder.
func (sapc *StudentAcademicProfileCreate) Mutation() *StudentAcademicProfileMutation {
	return sapc.mutation
}

// Save creates the StudentAcademicProfile in the database.
func (sapc *StudentAcademicProfileCreate) Save(ctx context.Context) (*StudentAcademicProfile, error) {
	var (
		err  error
		node *StudentAcademicProfile
	)
	sapc.defaults()
	if len(sapc.hooks) == 0 {
		if err = sapc.check(); err != nil {
			return nil, err
		}
		node, err = sapc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentAcademicProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sapc.check(); err != nil {
				return nil, err
			}
			sapc.mutation = mutation
			if node, err = sapc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sapc.hooks) - 1; i >= 0; i-- {
			if sapc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sapc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sapc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sapc *StudentAcademicProfileCreate) SaveX(ctx context.Context) *StudentAcademicProfile {
	v, err := sapc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sapc *StudentAcademicProfileCreate) Exec(ctx context.Context) error {
	_, err := sapc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sapc *StudentAcademicProfileCreate) ExecX(ctx context.Context) {
	if err := sapc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sapc *StudentAcademicProfileCreate) defaults() {
	if _, ok := sapc.mutation.Department(); !ok {
		v := studentacademicprofile.DefaultDepartment
		sapc.mutation.SetDepartment(v)
	}
	if _, ok := sapc.mutation.ID(); !ok {
		v := studentacademicprofile.DefaultID()
		sapc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sapc *StudentAcademicProfileCreate) check() error {
	if _, ok := sapc.mutation.Course(); !ok {
		return &ValidationError{Name: "course", err: errors.New(`ent: missing required field "StudentAcademicProfile.course"`)}
	}
	if v, ok := sapc.mutation.Course(); ok {
		if err := studentacademicprofile.CourseValidator(v); err != nil {
			return &ValidationError{Name: "course", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.course": %w`, err)}
		}
	}
	if _, ok := sapc.mutation.Institute(); !ok {
		return &ValidationError{Name: "institute", err: errors.New(`ent: missing required field "StudentAcademicProfile.institute"`)}
	}
	if v, ok := sapc.mutation.Institute(); ok {
		if err := studentacademicprofile.InstituteValidator(v); err != nil {
			return &ValidationError{Name: "institute", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.institute": %w`, err)}
		}
	}
	if _, ok := sapc.mutation.Board(); !ok {
		return &ValidationError{Name: "board", err: errors.New(`ent: missing required field "StudentAcademicProfile.board"`)}
	}
	if v, ok := sapc.mutation.Board(); ok {
		if err := studentacademicprofile.BoardValidator(v); err != nil {
			return &ValidationError{Name: "board", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.board": %w`, err)}
		}
	}
	if _, ok := sapc.mutation.RegNo(); !ok {
		return &ValidationError{Name: "reg_no", err: errors.New(`ent: missing required field "StudentAcademicProfile.reg_no"`)}
	}
	if v, ok := sapc.mutation.RegNo(); ok {
		if err := studentacademicprofile.RegNoValidator(v); err != nil {
			return &ValidationError{Name: "reg_no", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.reg_no": %w`, err)}
		}
	}
	if _, ok := sapc.mutation.Department(); !ok {
		return &ValidationError{Name: "department", err: errors.New(`ent: missing required field "StudentAcademicProfile.department"`)}
	}
	if _, ok := sapc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "StudentAcademicProfile.start_date"`)}
	}
	if v, ok := sapc.mutation.StartDate(); ok {
		if err := studentacademicprofile.StartDateValidator(v); err != nil {
			return &ValidationError{Name: "start_date", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.start_date": %w`, err)}
		}
	}
	if _, ok := sapc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "end_date", err: errors.New(`ent: missing required field "StudentAcademicProfile.end_date"`)}
	}
	if v, ok := sapc.mutation.EndDate(); ok {
		if err := studentacademicprofile.EndDateValidator(v); err != nil {
			return &ValidationError{Name: "end_date", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.end_date": %w`, err)}
		}
	}
	if _, ok := sapc.mutation.Marks(); !ok {
		return &ValidationError{Name: "marks", err: errors.New(`ent: missing required field "StudentAcademicProfile.marks"`)}
	}
	if v, ok := sapc.mutation.Marks(); ok {
		if err := studentacademicprofile.MarksValidator(v); err != nil {
			return &ValidationError{Name: "marks", err: fmt.Errorf(`ent: validator failed for field "StudentAcademicProfile.marks": %w`, err)}
		}
	}
	return nil
}

func (sapc *StudentAcademicProfileCreate) sqlSave(ctx context.Context) (*StudentAcademicProfile, error) {
	_node, _spec := sapc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sapc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected StudentAcademicProfile.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (sapc *StudentAcademicProfileCreate) createSpec() (*StudentAcademicProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &StudentAcademicProfile{config: sapc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: studentacademicprofile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentacademicprofile.FieldID,
			},
		}
	)
	if id, ok := sapc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sapc.mutation.Course(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldCourse,
		})
		_node.Course = value
	}
	if value, ok := sapc.mutation.Institute(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldInstitute,
		})
		_node.Institute = value
	}
	if value, ok := sapc.mutation.Board(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldBoard,
		})
		_node.Board = value
	}
	if value, ok := sapc.mutation.RegNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldRegNo,
		})
		_node.RegNo = value
	}
	if value, ok := sapc.mutation.Department(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldDepartment,
		})
		_node.Department = value
	}
	if value, ok := sapc.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldStartDate,
		})
		_node.StartDate = value
	}
	if value, ok := sapc.mutation.EndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldEndDate,
		})
		_node.EndDate = value
	}
	if value, ok := sapc.mutation.Marks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentacademicprofile.FieldMarks,
		})
		_node.Marks = value
	}
	if nodes := sapc.mutation.StudentProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentacademicprofile.StudentProfileTable,
			Columns: []string{studentacademicprofile.StudentProfileColumn},
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
		_node.student_profile_academic_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudentAcademicProfileCreateBulk is the builder for creating many StudentAcademicProfile entities in bulk.
type StudentAcademicProfileCreateBulk struct {
	config
	builders []*StudentAcademicProfileCreate
}

// Save creates the StudentAcademicProfile entities in the database.
func (sapcb *StudentAcademicProfileCreateBulk) Save(ctx context.Context) ([]*StudentAcademicProfile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sapcb.builders))
	nodes := make([]*StudentAcademicProfile, len(sapcb.builders))
	mutators := make([]Mutator, len(sapcb.builders))
	for i := range sapcb.builders {
		func(i int, root context.Context) {
			builder := sapcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentAcademicProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sapcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sapcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, sapcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sapcb *StudentAcademicProfileCreateBulk) SaveX(ctx context.Context) []*StudentAcademicProfile {
	v, err := sapcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sapcb *StudentAcademicProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := sapcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sapcb *StudentAcademicProfileCreateBulk) ExecX(ctx context.Context) {
	if err := sapcb.Exec(ctx); err != nil {
		panic(err)
	}
}