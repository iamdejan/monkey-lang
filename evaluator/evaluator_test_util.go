package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("wrong obj type. expected=`*object.Integer`, actual=`%T`", obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("wrong result.Value. expected=`%d`, actual=%d", expected, result.Value)
		return false
	}

	return true
}
