package setvar

import (
	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/nodes/expr"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type SetVar struct {
	basenode.BaseNode
	Params SetVarParams `mapstructure:"params"`
}

type SetVarParams struct {
	Name     string     `mapstructure:"name"`
	Type     string     `mapstructure:"type"`
	Readonly bool       `mapstructure:"readonly"`
	Value    expr.Value `mapstructure:"value"`
}

func (n *SetVar) Execute(vp *basevarpool.IVarPool) error {
	_v, err := n.Params.Value.Get(vp)
	if err != nil {
		return err
	}
	v := basevarpool.BaseVar{
		Value:    _v,
		Readonly: n.Params.Readonly,
	}
	return (*vp).SetVar(n.Params.Name, v)
}
