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
	"github.com/dlukt/graphql-backend-starter/ent/predicate"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks     []Hook
	mutation  *ProfileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *ProfileUpdate) SetUpdateTime(t time.Time) *ProfileUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (pu *ProfileUpdate) ClearUpdateTime() *ProfileUpdate {
	pu.mutation.ClearUpdateTime()
	return pu
}

// SetSub sets the "sub" field.
func (pu *ProfileUpdate) SetSub(s string) *ProfileUpdate {
	pu.mutation.SetSub(s)
	return pu
}

// SetNillableSub sets the "sub" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableSub(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetSub(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *ProfileUpdate) SetName(s string) *ProfileUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableName(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetGender sets the "gender" field.
func (pu *ProfileUpdate) SetGender(s string) *ProfileUpdate {
	pu.mutation.SetGender(s)
	return pu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableGender(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetGender(*s)
	}
	return pu
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProfileUpdate) defaults() error {
	if _, ok := pu.mutation.UpdateTime(); !ok && !pu.mutation.UpdateTimeCleared() {
		if profile.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized profile.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := profile.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfileUpdate) check() error {
	if v, ok := pu.mutation.Sub(); ok {
		if err := profile.SubValidator(v); err != nil {
			return &ValidationError{Name: "sub", err: fmt.Errorf(`ent: validator failed for field "Profile.sub": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *ProfileUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfileUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(profile.FieldUpdateTime, field.TypeTime, value)
	}
	if pu.mutation.UpdateTimeCleared() {
		_spec.ClearField(profile.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := pu.mutation.Sub(); ok {
		_spec.SetField(profile.FieldSub, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Gender(); ok {
		_spec.SetField(profile.FieldGender, field.TypeString, value)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProfileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (puo *ProfileUpdateOne) SetUpdateTime(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (puo *ProfileUpdateOne) ClearUpdateTime() *ProfileUpdateOne {
	puo.mutation.ClearUpdateTime()
	return puo
}

// SetSub sets the "sub" field.
func (puo *ProfileUpdateOne) SetSub(s string) *ProfileUpdateOne {
	puo.mutation.SetSub(s)
	return puo
}

// SetNillableSub sets the "sub" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableSub(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetSub(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *ProfileUpdateOne) SetName(s string) *ProfileUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableName(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetGender sets the "gender" field.
func (puo *ProfileUpdateOne) SetGender(s string) *ProfileUpdateOne {
	puo.mutation.SetGender(s)
	return puo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableGender(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetGender(*s)
	}
	return puo
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProfileUpdate builder.
func (puo *ProfileUpdateOne) Where(ps ...predicate.Profile) *ProfileUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProfileUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdateTime(); !ok && !puo.mutation.UpdateTimeCleared() {
		if profile.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized profile.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := profile.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfileUpdateOne) check() error {
	if v, ok := puo.mutation.Sub(); ok {
		if err := profile.SubValidator(v); err != nil {
			return &ValidationError{Name: "sub", err: fmt.Errorf(`ent: validator failed for field "Profile.sub": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *ProfileUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfileUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Profile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(profile.FieldUpdateTime, field.TypeTime, value)
	}
	if puo.mutation.UpdateTimeCleared() {
		_spec.ClearField(profile.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := puo.mutation.Sub(); ok {
		_spec.SetField(profile.FieldSub, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Gender(); ok {
		_spec.SetField(profile.FieldGender, field.TypeString, value)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
