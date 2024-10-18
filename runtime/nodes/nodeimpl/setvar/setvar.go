package setvar

import (
	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type SetVar struct {
	basenode.BaseNode
	Params SetVarParams `json:"params"`
}

type SetVarParams struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Readonly bool        `json:"readonly"`
	Value    interface{} `json:"value"`
}

func (n *SetVar) Execute(vp *basevarpool.IVarPool) error {
	v := basevarpool.BaseVar{
		Value:    n.Params.Value,
		Readonly: n.Params.Readonly,
	}
	return (*vp).SetVar(n.Params.Name, v)
}
