package runtime

import (
	"encoding/json"

	"github.com/yiipu/i9r/runtime/graphengine"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type Runtime struct {
	Engine  graphengine.GraphEngine
	VarPool basevarpool.IVarPool
}

func (r *Runtime) Execute(data []byte) error {
	var graph graphengine.Graph
	err := json.Unmarshal(data, &graph)
	if err != nil {
		return err
	}

	r.VarPool.Init()

	r.Engine.Graph = &graph
	r.Engine.VarPool = r.VarPool

	err = r.Engine.Execute()
	if err != nil {
		return err
	}

	return nil
}
