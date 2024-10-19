package runtime

import (
	"os"
	"testing"

	"github.com/yiipu/i9r/runtime/graphengine"
	"github.com/yiipu/i9r/runtime/varpool/memvarpool"
)

func TestRuntime_Execute(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "graph with set eval if",
			filePath: "./.testdata/valid/graph.json",
			wantErr:  false,
		},
		{
			name:     "simple graph with set",
			filePath: "./.testdata/valid/simple.json",
			wantErr:  false,
		},
		{
			name:     "increment",
			filePath: "./.testdata/valid/increment.json",
			wantErr:  false,
		},
		{
			name:     "calc sum 1 to 10",
			filePath: "./.testdata/valid/calcSum1to10.json",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.filePath)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}

			r := &Runtime{
				Engine: graphengine.GraphEngine{VarPool: &memvarpool.MemVarPool{}},
			}
			err = r.Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("Runtime.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			err = r.Execute(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Runtime.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
