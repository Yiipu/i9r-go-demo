package runtime

import (
	"encoding/json"

	"github.com/yiipu/i9r/runtime/graphengine"
)

type Runtime struct {
	Engine graphengine.GraphEngine
}

func (r *Runtime) Init() error {
	err := r.Engine.VarPool.Init()
	if err != nil {
		return err
	}

	return nil
}

func (r *Runtime) Execute(data []byte) error {
	var graph graphengine.Graph
	err := json.Unmarshal(data, &graph)
	if err != nil {
		return err
	}

	r.Engine.VarPool.Clear()
	r.Engine.Graph = &graph

	err = r.Engine.Execute()
	if err != nil {
		return err
	}

	return nil
}
