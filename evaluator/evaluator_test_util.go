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

func testIntegerObject(t *testing.T, obj object.Object, expected int64, input string) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("wrong obj type for input `%s`. expected=`*object.Integer`, actual=`%T`", input, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("wrong result.Value for input `%s`. expected=%d, actual=%d", input, expected, result.Value)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool, input string) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("wrong obj type for input `%s`. expected=`*object.Integer`, actual=`%T`", input, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("wrong result.Value for input `%s`. expected=`%t`, actual=%t", input, expected, result.Value)
		return false
	}

	return true
}
