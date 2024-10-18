// basevarpool provides the interface for variable pool.
package basevarpool

type BaseVar struct {
	Value    interface{}
	Readonly bool
}

// MagicVar is a special variable that is bound to a source.
type MagicVar struct {
	BaseVar
}

type IVarPool interface {
	Init() error
	SetVar(name string, v BaseVar) error
	GetVar(name string) (BaseVar, error)
	SetMagicVar(source string, v MagicVar) error
	GetMagicVar(source string) (MagicVar, error)
}
