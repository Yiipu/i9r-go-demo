package evaluate

import (
	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/nodes/expr"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type Evaluate struct {
	basenode.BaseNode
	Params EvaluateParams `json:"params"`
}

type EvaluateParams struct {
	Expression expr.AlgoExpr `json:"expression"`
}

func (n *Evaluate) Execute(vp *basevarpool.IVarPool) error {
	result, err := n.Params.Expression.Evaluate(vp)
	if err != nil {
		return err
	}

	(*vp).SetMagicVar(n.ID, basevarpool.MagicVar{
		BaseVar: basevarpool.BaseVar{
			Value:    result,
			Readonly: true,
		},
	})

	return nil
}
