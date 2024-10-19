package mapping

import (
	"github.com/mitchellh/mapstructure"

	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/nodes/nodeimpl/evaluate"
	"github.com/yiipu/i9r/runtime/nodes/nodeimpl/format"
	"github.com/yiipu/i9r/runtime/nodes/nodeimpl/ifelse"
	"github.com/yiipu/i9r/runtime/nodes/nodeimpl/output"
	"github.com/yiipu/i9r/runtime/nodes/nodeimpl/setvar"
)

const (
	IF       basenode.NodeType = "if"
	EVALUATE basenode.NodeType = "evaluate"
	SETVAR   basenode.NodeType = "setVariable"
	LOG      basenode.NodeType = "log"
	FORMAT   basenode.NodeType = "format"
)

var Mapping = map[basenode.NodeType]func(basenode.BaseNode) basenode.NodeInterface{
	IF: func(b basenode.BaseNode) basenode.NodeInterface {
		n := &ifelse.IfElse{BaseNode: b}
		mapstructure.Decode(b.Params, &n.Params)
		return n
	},
	EVALUATE: func(b basenode.BaseNode) basenode.NodeInterface {
		n := &evaluate.Evaluate{BaseNode: b}
		mapstructure.Decode(b.Params, &n.Params)
		return n
	},
	SETVAR: func(b basenode.BaseNode) basenode.NodeInterface {
		n := &setvar.SetVar{BaseNode: b}
		mapstructure.Decode(b.Params, &n.Params)
		return n
	},
	LOG: func(b basenode.BaseNode) basenode.NodeInterface {
		n := &output.Output{BaseNode: b}
		mapstructure.Decode(b.Params, &n.Params)
		return n
	},
	FORMAT: func(b basenode.BaseNode) basenode.NodeInterface {
		n := &format.Format{BaseNode: b}
		mapstructure.Decode(b.Params, &n.Params)
		return n
	},
}
