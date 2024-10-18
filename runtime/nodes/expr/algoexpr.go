package expr

import (
	"errors"

	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type algoOpType string

const (
	Add algoOpType = "add"
	Sub algoOpType = "sub"
)

type AlgoExpr struct {
	Operator algoOpType `json:"operator"`
	Left     Value      `json:"left"`
	Right    Value      `json:"right"`
}

func (a *AlgoExpr) Evaluate(vp *basevarpool.IVarPool) (interface{}, error) {
	l, err := a.Left.Get(vp)
	if err != nil {
		return nil, err
	}
	r, err := a.Right.Get(vp)
	if err != nil {
		return nil, err
	}

	switch a.Operator {
	case Add:
		return l.(float64) + r.(float64), nil
	case Sub:
		return l.(float64) - r.(float64), nil
	}
	return nil, errors.New("unknown operator")
}
