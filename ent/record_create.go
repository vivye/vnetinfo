// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"netinfo/ent/record"
	"netinfo/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordCreate is the builder for creating a Record entity.
type RecordCreate struct {
	config
	mutation *RecordMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *RecordCreate) SetCreatedAt(t time.Time) *RecordCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RecordCreate) SetNillableCreatedAt(t *time.Time) *RecordCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RecordCreate) SetUpdatedAt(t time.Time) *RecordCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RecordCreate) SetNillableUpdatedAt(t *time.Time) *RecordCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetDescription sets the "description" field.
func (rc *RecordCreate) SetDescription(s string) *RecordCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNetInterfaces sets the "net_interfaces" field.
func (rc *RecordCreate) SetNetInterfaces(si []schema.NetInterface) *RecordCreate {
	rc.mutation.SetNetInterfaces(si)
	return rc
}

// SetID sets the "id" field.
func (rc *RecordCreate) SetID(u uint) *RecordCreate {
	rc.mutation.SetID(u)
	return rc
}

// Mutation returns the RecordMutation object of the builder.
func (rc *RecordCreate) Mutation() *RecordMutation {
	return rc.mutation
}

// Save creates the Record in the database.
func (rc *RecordCreate) Save(ctx context.Context) (*Record, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecordCreate) SaveX(ctx context.Context) *Record {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecordCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecordCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RecordCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := record.DefaultCreatedAt
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := record.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecordCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Record.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Record.updated_at"`)}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Record.description"`)}
	}
	if _, ok := rc.mutation.NetInterfaces(); !ok {
		return &ValidationError{Name: "net_interfaces", err: errors.New(`ent: missing required field "Record.net_interfaces"`)}
	}
	return nil
}

func (rc *RecordCreate) sqlSave(ctx context.Context) (*Record, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RecordCreate) createSpec() (*Record, *sqlgraph.CreateSpec) {
	var (
		_node = &Record{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(record.Table, sqlgraph.NewFieldSpec(record.FieldID, field.TypeUint))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(record.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(record.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(record.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.NetInterfaces(); ok {
		_spec.SetField(record.FieldNetInterfaces, field.TypeJSON, value)
		_node.NetInterfaces = value
	}
	return _node, _spec
}

// RecordCreateBulk is the builder for creating many Record entities in bulk.
type RecordCreateBulk struct {
	config
	builders []*RecordCreate
}

// Save creates the Record entities in the database.
func (rcb *RecordCreateBulk) Save(ctx context.Context) ([]*Record, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Record, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecordMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
					nodes[i].ID = uint(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecordCreateBulk) SaveX(ctx context.Context) []*Record {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecordCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecordCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}