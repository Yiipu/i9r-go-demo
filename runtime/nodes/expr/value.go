package expr

import (
	"errors"

	"github.com/mitchellh/mapstructure"

	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type StorageType string
type ValueType string

const (
	MagicVariable StorageType = "magicVariable"
	Variable      StorageType = "variable"
	Constant      StorageType = "constant"
)
const (
	Number ValueType = "number"
	String ValueType = "string"
)

type Value struct {
	Type StorageType            `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type VariableData struct {
	Name string `json:"name"`
}
type MagicVariableData struct {
	Source string `json:"source"`
}
type ConstantData struct {
	Type  ValueType   `json:"type"`
	Value interface{} `json:"value"`
}

func (v *Value) Get(vp *basevarpool.IVarPool) (interface{}, error) {
	switch v.Type {
	case MagicVariable:
		data := MagicVariableData{}
		err := mapstructure.Decode(v.Data, &data)
		if err != nil {
			return nil, err
		}
		magicVar, err := (*vp).GetMagicVar(data.Source)
		if err != nil {
			return nil, err
		}
		return magicVar.Value, nil
	case Variable:
		data := VariableData{}
		err := mapstructure.Decode(v.Data, &data)
		if err != nil {
			return nil, err
		}
		variable, err := (*vp).GetVar(data.Name)
		if err != nil {
			return nil, err
		}
		return variable.Value, nil
	case Constant:
		data := ConstantData{}
		err := mapstructure.Decode(v.Data, &data)
		if err != nil {
			return nil, err
		}
		return data.Value, nil
	}
	return nil, errors.New("unknown storage type")
}
