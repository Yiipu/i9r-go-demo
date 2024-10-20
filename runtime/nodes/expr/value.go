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
	DataStorageType StorageType            `mapstructure:"type"`
	Data            map[string]interface{} `mapstructure:"data"`
}

type VariableData struct {
	Name string `mapstructure:"name"`
}
type MagicVariableData struct {
	Source string `mapstructure:"source"`
}
type ConstantData struct {
	Type  ValueType   `mapstructure:"type"`
	Value interface{} `mapstructure:"value"`
}

func (v *Value) Get(vp *basevarpool.IVarPool) (interface{}, error) {
	switch v.DataStorageType {
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
