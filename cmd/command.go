package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yiipu/i9r/runtime"
	"github.com/yiipu/i9r/runtime/graphengine"
	"github.com/yiipu/i9r/runtime/varpool/memvarpool"
)

func main() {
	filePath := flag.String("file-path", "", "Path to the JSON file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a file path using --file-path")
		os.Exit(1)
	}

	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	r := &runtime.Runtime{
		Engine: graphengine.GraphEngine{VarPool: &memvarpool.MemVarPool{}},
	}

	err = r.Init()
	if err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		os.Exit(1)
	}

	err = r.Execute(data)
	if err != nil {
		fmt.Printf("Execution failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Execution succeeded")
}
