package output

import (
	"fmt"

	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/nodes/expr"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type Output struct {
	basenode.BaseNode
	Params OutputParams `mapstructure:"params"`
}

type OutputParams struct {
	Message expr.Value `mapstructure:"message"`
	Target  string     `mapstructure:"target"`
}

func (n *Output) Execute(vp *basevarpool.IVarPool) error {
	// only console for now
	if n.Params.Target == "console" {
		message, err := n.Params.Message.Get(vp)
		if err != nil {
			return err
		}
		fmt.Println(message)
	}
	return nil
}
