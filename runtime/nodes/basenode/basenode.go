package basenode

import (
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type NodeType string

type BaseNode struct {
	ID     string                 `json:"id"`
	Type   NodeType               `json:"type"`
	Params map[string]interface{} `json:"parameters"`
}

type NodeInterface interface {
	Execute(vp *basevarpool.IVarPool) error
}

func (n *BaseNode) Execute(vp *basevarpool.IVarPool) error {
	panic("not implemented")
}
