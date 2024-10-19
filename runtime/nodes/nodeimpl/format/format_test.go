package format

import (
	"testing"

	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
	"github.com/yiipu/i9r/runtime/varpool/memvarpool"
)

func TestFormat_Execute(t *testing.T) {
	tests := []struct {
		name         string
		formatString string
		vars         []basevarpool.BaseVar
		varnames     []string
		mvars        []basevarpool.MagicVar
		mvarnames    []string
		wanted       string
	}{
		{
			"str", "hello, {{v:var1}}",
			[]basevarpool.BaseVar{{Value: "world"}}, []string{"var1"},
			[]basevarpool.MagicVar{}, []string{},
			"hello, world",
		},
		{
			"int", "hello, {{v:var1}}",
			[]basevarpool.BaseVar{{Value: 123}}, []string{"var1"},
			[]basevarpool.MagicVar{}, []string{},
			"hello, 123",
		},
		{
			"float", "hello, {{v:var1}}",
			[]basevarpool.BaseVar{{Value: 123.456}}, []string{"var1"},
			[]basevarpool.MagicVar{}, []string{},
			"hello, 123.456",
		},
		{
			"str-in-magic", "hello, {{m:var1}}",
			[]basevarpool.BaseVar{}, []string{},
			[]basevarpool.MagicVar{{BaseVar: basevarpool.BaseVar{Value: "world"}}}, []string{"var1"},
			"hello, world",
		},
		{
			"combined", "hello, {{v:var1}} and {{m:var2}}, {{v:var2}} and {{m:var1}}",
			[]basevarpool.BaseVar{{Value: "world"}, {Value: 123}}, []string{"var1", "var2"},
			[]basevarpool.MagicVar{{BaseVar: basevarpool.BaseVar{Value: "universe"}}, {BaseVar: basevarpool.BaseVar{Value: true}}}, []string{"var2", "var1"},
			"hello, world and universe, 123 and true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vp := setupVarPool(t, tt.vars, tt.varnames, tt.mvars, tt.mvarnames)
			n := createFormatNode(tt.name, tt.formatString)
			executeAndVerify(t, n, vp, tt.wanted)
		})
	}
}

func setupVarPool(t *testing.T, vars []basevarpool.BaseVar, varnames []string, mvars []basevarpool.MagicVar, mvarnames []string) *memvarpool.MemVarPool {
	vp := &memvarpool.MemVarPool{}
	if err := vp.Init(); err != nil {
		t.Fatalf("MemVarPool.Init() error = %v", err)
	}
	for k, v := range vars {
		if err := vp.SetVar(varnames[k], v); err != nil {
			t.Fatalf("MemVarPool.SetVar() error = %v", err)
		}
	}
	for k, v := range mvars {
		if err := vp.SetMagicVar(mvarnames[k], v); err != nil {
			t.Fatalf("MemVarPool.SetMagicVar() error = %v", err)
		}
	}
	return vp
}

func createFormatNode(name, formatString string) *Format {
	return &Format{
		BaseNode: basenode.BaseNode{
			ID:   name,
			Type: "format",
		},
		Params: FormatParams{
			FormatString: formatString,
		},
	}
}

func executeAndVerify(t *testing.T, n *Format, vp *memvarpool.MemVarPool, wanted string) {
	var ivp basevarpool.IVarPool = vp
	if err := n.Execute(&ivp); err != nil {
		t.Fatalf("Format.Execute() error = %v", err)
	}
	output, err := vp.GetMagicVar(n.ID)
	if err != nil {
		t.Fatalf("MemVarPool.GetMagicVar() error = %v", err)
	}
	if output.Value != wanted {
		t.Fatalf("Format.Execute() = %v, want %v", output.Value, wanted)
	}
}
