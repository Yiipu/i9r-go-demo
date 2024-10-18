package graphengine

import (
	"github.com/yiipu/i9r/runtime/nodes/basenode"
)

type Graph struct {
	Nodes map[string]basenode.BaseNode `json:"nodes"`
	Edges map[string][]Target          `json:"edges"`
}

type Target struct {
	Target    string `json:"target"`
	Condition string `json:"condition"`
}
