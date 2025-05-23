// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
	"github.com/rs/xid"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *ProfileCreate) SetCreateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProfileCreate) SetUpdateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetSub sets the "sub" field.
func (pc *ProfileCreate) SetSub(s string) *ProfileCreate {
	pc.mutation.SetSub(s)
	return pc
}

// SetName sets the "name" field.
func (pc *ProfileCreate) SetName(s string) *ProfileCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetGender sets the "gender" field.
func (pc *ProfileCreate) SetGender(s string) *ProfileCreate {
	pc.mutation.SetGender(s)
	return pc
}

// SetID sets the "id" field.
func (pc *ProfileCreate) SetID(x xid.ID) *ProfileCreate {
	pc.mutation.SetID(x)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableID(x *xid.ID) *ProfileCreate {
	if x != nil {
		pc.SetID(*x)
	}
	return pc
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		if profile.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized profile.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := profile.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if profile.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized profile.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := profile.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		if profile.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized profile.DefaultID (forgotten import ent/runtime?)")
		}
		v := profile.DefaultID()
		pc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Profile.create_time"`)}
	}
	if _, ok := pc.mutation.Sub(); !ok {
		return &ValidationError{Name: "sub", err: errors.New(`ent: missing required field "Profile.sub"`)}
	}
	if v, ok := pc.mutation.Sub(); ok {
		if err := profile.SubValidator(v); err != nil {
			return &ValidationError{Name: "sub", err: fmt.Errorf(`ent: validator failed for field "Profile.sub": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Profile.name"`)}
	}
	if _, ok := pc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Profile.gender"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(profile.Table, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(profile.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(profile.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Sub(); ok {
		_spec.SetField(profile.FieldSub, field.TypeString, value)
		_node.Sub = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.SetField(profile.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	return _node, _spec
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	err      error
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
