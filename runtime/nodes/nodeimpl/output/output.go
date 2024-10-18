package output

import (
	"fmt"

	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type Output struct {
	basenode.BaseNode
	Params OutputParams `json:"params"`
}

type OutputParams struct {
	Message string `json:"message"`
	Target  string `json:"target"`
}

func (n *Output) Execute(vp *basevarpool.IVarPool) error {
	if n.Params.Target == "console" {
		fmt.Println(n.Params.Message)
	}
	return nil
}
