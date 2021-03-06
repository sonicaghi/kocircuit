package macros

import (
	"fmt"
	"reflect"

	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
	. "github.com/kocircuit/kocircuit/lang/go/eval/symbol"
	. "github.com/kocircuit/kocircuit/lang/go/kit/tree"
	. "github.com/kocircuit/kocircuit/lang/go/kit/util"
)

type EvalGoValueMacro struct {
	Value interface{} `ko:"name=value"`
}

func (m EvalGoValueMacro) MacroID() string { return m.Help() }

func (m EvalGoValueMacro) Label() string { return "constant" }

func (m EvalGoValueMacro) MacroSheathString() *string { return PtrString(m.Help()) }

func (m EvalGoValueMacro) Help() string {
	return fmt.Sprintf("Constant(%s)", Sprint(m.Value))
}

func (m EvalGoValueMacro) Doc() string {
	return fmt.Sprintf("Returns the go value %v", m.Value)
}

func (m EvalGoValueMacro) Invoke(span *Span, arg Arg) (returns Return, effect Effect, err error) {
	return Deconstruct(span, reflect.ValueOf(m.Value)), nil, nil
}

type EvalSymbolMacro struct {
	Symbol Symbol `ko:"name=symbol"`
}

func (m EvalSymbolMacro) MacroID() string { return m.Help() }

func (m EvalSymbolMacro) Label() string { return "symbol" }

func (m EvalSymbolMacro) MacroSheathString() *string { return PtrString(m.Help()) }

func (m EvalSymbolMacro) Help() string {
	return fmt.Sprintf("Symbol(%s)", Sprint(m.Symbol))
}

func (m EvalSymbolMacro) Doc() string {
	return fmt.Sprintf("Returns the value %v", m.Symbol)
}

func (m EvalSymbolMacro) Invoke(span *Span, arg Arg) (returns Return, effect Effect, err error) {
	return m.Symbol, nil, nil
}
