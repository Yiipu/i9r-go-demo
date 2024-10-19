package setvar

import (
	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/nodes/expr"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type SetVar struct {
	basenode.BaseNode
	Params SetVarParams `json:"params"`
}

type SetVarParams struct {
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Readonly bool       `json:"readonly"`
	Value    expr.Value `json:"value"`
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
