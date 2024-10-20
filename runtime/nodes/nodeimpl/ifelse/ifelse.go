package ifelse

import (
	"errors"

	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/nodes/expr"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type IfElse struct {
	basenode.BaseNode
	Params IfElseParams `json:"params"`
}

type LogicOperator string

const (
	AND LogicOperator = "and"
	OR  LogicOperator = "or"
)

type IfElseParams struct {
	Operator   LogicOperator   `json:"operator"`
	Conditions []expr.BoolExpr `json:"conditions"`
}

func (n *IfElse) Execute(vp *basevarpool.IVarPool) error {
	var result bool
	var err error

	switch n.Params.Operator {
	case AND:
		result, err = evaluateConditions(vp, n.Params.Conditions, false)
	case OR:
		result, err = evaluateConditions(vp, n.Params.Conditions, true)
	default:
		return errors.New("unknown operator")
	}

	if err != nil {
		return err
	}

	err = (*vp).SetMagicVar(n.ID, basevarpool.MagicVar{
		BaseVar: basevarpool.BaseVar{
			Value:    result,
			Readonly: false,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func evaluateConditions(vp *basevarpool.IVarPool, conditions []expr.BoolExpr, shortCircuit bool) (bool, error) {
	for _, c := range conditions {
		r, err := c.Evaluate(vp)
		if err != nil {
			return false, err
		}
		if r == shortCircuit {
			return shortCircuit, nil
		}
	}
	return !shortCircuit, nil
}
