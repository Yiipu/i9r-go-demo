// mvm is a memory based var manager implementation.
package memvarpool

import (
	"errors"

	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type pool struct {
	vars     map[string]basevarpool.BaseVar
	capacity int
	size     int
}

type MemVarPool struct {
	varpool      pool
	magicvarpool pool
}

func (m *MemVarPool) Init() error {
	// TODO: configure capacity
	m.varpool = pool{
		vars:     make(map[string]basevarpool.BaseVar),
		capacity: 100,
		size:     0,
	}
	m.magicvarpool = pool{
		vars:     make(map[string]basevarpool.BaseVar),
		capacity: 100,
		size:     0,
	}
	return nil
}

func (m *MemVarPool) setVar(p pool, n string, v basevarpool.BaseVar) error {
	if p.size >= p.capacity {
		return errors.New("pool is full")
	}
	if existingVar, ok := p.vars[n]; ok && existingVar.Readonly {
		return errors.New("var is readonly")
	}

	p.vars[n] = v
	p.size++
	return nil
}

func (m *MemVarPool) getVar(p pool, n string) (basevarpool.BaseVar, error) {
	v, ok := p.vars[n]
	if !ok {
		return basevarpool.BaseVar{}, errors.New("var not found")
	}
	return v, nil
}

func (m *MemVarPool) SetVar(source string, v basevarpool.BaseVar) error {
	return m.setVar(m.varpool, source, v)
}

func (m *MemVarPool) GetVar(source string) (basevarpool.BaseVar, error) {
	return m.getVar(m.varpool, source)
}

func (m *MemVarPool) SetMagicVar(source string, v basevarpool.MagicVar) error {
	_v := basevarpool.BaseVar{
		Value:    v.Value,
		Readonly: v.Readonly,
	}
	return m.setVar(m.magicvarpool, source, _v)
}

func (m *MemVarPool) GetMagicVar(source string) (basevarpool.MagicVar, error) {
	v, err := m.getVar(m.magicvarpool, source)
	if err != nil {
		return basevarpool.MagicVar{}, err
	}
	return basevarpool.MagicVar{
		BaseVar: v,
	}, nil
}
