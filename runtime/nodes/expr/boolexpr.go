package expr

import (
	"errors"

	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type boolOpType string

const (
	// for number
	Eq  boolOpType = "eq"
	Neq boolOpType = "neq"
	Geq boolOpType = "geq"
	Leq boolOpType = "leq"
	Gt  boolOpType = "gt"
	Lt  boolOpType = "lt"

	// for string or array
	Contains    boolOpType = "contains"
	NotContains boolOpType = "not contains"
)

type BoolExpr struct {
	Operator boolOpType `mapstructure:"operator"`
	Left     Value      `mapstructure:"left"`
	Right    Value      `mapstructure:"right"`
}

func (b *BoolExpr) Evaluate(vp *basevarpool.IVarPool) (bool, error) {
	l, err := b.Left.Get(vp)
	if err != nil {
		return false, err
	}
	r, err := b.Right.Get(vp)
	if err != nil {
		return false, err
	}

	switch b.Operator {
	// TODO: handle assert panic
	case Eq:
		return l == r, nil
	case Neq:
		return l != r, nil
	case Geq:
		return l.(float64) >= r.(float64), nil
	case Leq:
		return l.(float64) <= r.(float64), nil
	case Gt:
		return l.(float64) > r.(float64), nil
	case Lt:
		return l.(float64) < r.(float64), nil
	case Contains:
		panic("not implemented")
	case NotContains:
		panic("not implemented")
	}
	return false, errors.New("unknown operator")
}
