package graphengine

import (
	"errors"

	"github.com/yiipu/i9r/runtime/nodes/mapping"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type GraphEngine struct {
	VarPool basevarpool.IVarPool
	Graph   *Graph
}

func (ge *GraphEngine) Execute() error {
	node := ge.Graph.Nodes["start"]
outer_loop:
	for {
		nodeType := node.Type
		nodeFactory, exists := mapping.Mapping[nodeType]
		if !exists {
			return errors.New("node type not found")
		}
		nodeImpl := nodeFactory(node)
		err := nodeImpl.Execute(&ge.VarPool)
		if err != nil {
			return err
		}
		// TODO: improve contitional edge handling
		lastoutput, err := ge.VarPool.GetMagicVar(string(node.ID))
		for _, target := range ge.Graph.Edges[node.ID] {
			if target.Condition == "" {
				node = ge.Graph.Nodes[target.Target]
				continue outer_loop
			}
			booloutput := lastoutput.Value.(bool)
			condition := target.Condition
			if err != nil {
				return errors.New("encountered contitional edge with out conditional node")
			}
			switch condition {
			case "true":
				if booloutput {
					node = ge.Graph.Nodes[target.Target]
					continue outer_loop
				}
			case "false":
				if !booloutput {
					node = ge.Graph.Nodes[target.Target]
					continue outer_loop
				}
			}
		}
		break
	}
	return nil
}
