package macros

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
	. "github.com/kocircuit/kocircuit/lang/go/eval"
	. "github.com/kocircuit/kocircuit/lang/go/eval/symbol"
	. "github.com/kocircuit/kocircuit/lang/go/kit/util"
)

func init() {
	RegisterEvalMacro("All", new(EvalAllMacro))
}

type EvalAllMacro struct{}

func (m EvalAllMacro) MacroID() string { return m.Help() }

func (m EvalAllMacro) Label() string { return "all" }

func (m EvalAllMacro) MacroSheathString() *string { return PtrString("All") }

func (m EvalAllMacro) Help() string { return "All" }

func (m EvalAllMacro) Doc() string {
	return `
The builtin function All() is designed to facilitate branching on
the condition that a few values are all non-empty.

All() accepts any number of named arguments.
If all arguments are non-empty, All() returns a structure containing all named arguments as fields.
If any one of the arguments is empty, All() returns the empty value.`
}

func (EvalAllMacro) Invoke(span *Span, arg Arg) (returns Return, effect Effect, err error) {
	a := arg.(*StructSymbol)
	all := true
	for _, field := range a.Field {
		if IsEmptySymbol(field.Value) {
			all = false
			break // for
		}
	}
	if all && len(a.Field) > 0 {
		return a, nil, nil
	} else {
		return EmptySymbol{}, nil, nil
	}
}
