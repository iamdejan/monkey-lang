package evaluator

import (
	"monkey/object"
	"testing"
)

type IntegerEvalTest struct {
	input    string
	expected int64
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []IntegerEvalTest{
		{input: "5", expected: 5},
		{input: "10", expected: 10},
		{input: "-5", expected: -5},
		{input: "-10", expected: -10},
		{input: "1+1", expected: 2},
		{input: "1 + 1", expected: 2},
		{input: "1 + 2 + 3", expected: 6},
		{input: "5 + 5 + 5 + 5 -10", expected: 10},
		{input: "2 * 2 * 2", expected: 8},
		{input: "-50 + 100 + -50", expected: 0},
		{input: "5 * 2 - 10", expected: 0},
		{input: "5 / 2", expected: 2}, // handle decimal division later
		{input: "2 * (5 + 3)", expected: 16},
		{input: "3 * (3 * 3) - 6", expected: 21},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected, tt.input)
	}
}

type BoolEvalTest struct {
	input    string
	expected bool
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []BoolEvalTest{
		{input: "true", expected: true},
		{input: "false", expected: false},
		{input: "1 < 2", expected: true},
		{input: "1 > 2", expected: false},
		{input: "1 == 1", expected: true},
		{input: "1 != 2", expected: true},
		{input: "1 == 2", expected: false},
		{input: "1+2 == 3", expected: true},
		{input: "2*4 != 2", expected: true},
		{input: "2*4 == 2", expected: false},
		{input: "(1 < 2) == true", expected: true},
		{input: "(1 < 2) == false", expected: false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected, tt.input)
	}
}

type BangOperatorTest struct {
	input    string
	expected bool
}

func TestBangOperator(t *testing.T) {
	tests := []BangOperatorTest{
		{input: "!true", expected: false},
		{input: "!false", expected: true},
		{input: "!5", expected: false},
		{input: "!!true", expected: true},
		{input: "!!false", expected: false},
		{input: "!!5", expected: true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected, tt.input)
	}
}

type IfElseTest struct {
	input    string
	expected interface{}
}

func TestIfElseExpression(t *testing.T) {
	tests := []IfElseTest{
		{input: "if (true) { 10 }", expected: 10},
		{input: "if (false) { 10 }", expected: nil},
		{input: "if (1) { 10 }", expected: 10},
		{input: "if (1 < 2) { 10 }", expected: 10},
		{input: "if (1 > 2) { 10 }", expected: nil},
		{input: "if (1 < 2) { 10 } else { 20 }", expected: 10},
		{input: "if (1 > 2) { 10 } else { 20 }", expected: 20},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer), tt.input)
		} else {
			testNullObject(t, evaluated, tt.input)
		}
	}
}

type ReturnTest struct {
	input    string
	expected interface{}
}

func TestReturnStatement(t *testing.T) {
	tests := []ReturnTest{
		{input: "return 5;", expected: 5},
		{input: "return -10;", expected: -10},
		{input: "return 1-2;", expected: 1 - 2},
		{input: "return 1-2+3;", expected: 1 - 2 + 3},
		{input: "return true;", expected: true},
		{input: "return false;", expected: false},
		{input: "return null;", expected: nil},
		{input: "5; return null; 5", expected: nil},
		{input: "return 10; 9;", expected: 10},
		{input: "return 2 * 5; 9;", expected: 2 * 5},
		{input: "9; return 2 * 5; 9;", expected: 2 * 5},
		{input: `
		if (10 > 1) {
			if (10 > 1) {
				return 10;
			}

			return 1;
		}`, expected: 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer), tt.input)
			return
		}

		boolean, ok := tt.expected.(bool)
		if ok {
			testBooleanObject(t, evaluated, boolean, tt.input)
			return
		}

		testNullObject(t, evaluated, tt.input)
	}
}

type ErrorTest struct {
	input           string
	expectedMessage string
}

func TestErrorHandling(t *testing.T) {
	tests := []ErrorTest{
		{input: "5 + true", expectedMessage: "type mismatch: INTEGER + BOOLEAN"},
		{input: "5 + true; 5;", expectedMessage: "type mismatch: INTEGER + BOOLEAN"},
		{input: "-true;", expectedMessage: "unknown operator: -BOOLEAN"},
		{input: "-true; 5;", expectedMessage: "unknown operator: -BOOLEAN"},
		{input: "true + false;", expectedMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{input: "5; true + false; 5;", expectedMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{input: "if(10 > 1) { true + false; }", expectedMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{input: `
		if (10 > 1) {
			if (10 > 1) {
				return true + false;
			}

			return 1;
		}`, expectedMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{input: "foobar;", expectedMessage: "identifier not found: foobar"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. actual=`%T(%#v)`", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=`%s`, actual=`%s`", tt.expectedMessage, errObj.Message)
		}
	}
}

type LetTest struct {
	input    string
	expected int64
}

func TestLetStatements(t *testing.T) {
	tests := []LetTest{
		{input: "let a = 5; a;", expected: 5},
		{input: "let a = 5 * 5; a;", expected: 5 * 5},
		{input: "let a = 5; let b = a; b;", expected: 5},
		{input: "let a = 5; let b = a; let c = a + b + 5; c;", expected: 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected, tt.input)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("`evaluated` object got wrong type. expected=`*object.Function`, actual=`%T`", evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function `fn` has wrong parameter count. expected=`1`, actual=`%d`", len(fn.Parameters))
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("wrong parameter name. expected=`x`, actual=`%s`", fn.Parameters[0].String())
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("wrong `fn.Body`. expected=`%s`, actual=`%s`", expectedBody, fn.Body.String())
	}
}

type FunctionCallTest struct {
	input string
	expected int64
}

func TestFunctionApplication(t *testing.T) {
	tests := []FunctionCallTest{
		{input: "let identity = fn(x) { x; }; identity(5);", expected: 5},
		{input: "let identity = fn(x) { return x; }; identity(5);", expected: 5},
		{input: "let double = fn(x) { return x * 2; }; double(5);", expected: 5 * 2},
		{input: "let add = fn(x, y) { x + y; }; add(1, 2);", expected: 1 + 2},
		{input: "let add = fn(x, y) { x + y; }; add(5, add(1, 2));", expected: 8},
		{input: "fn(x) { x; }(5);", expected: 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected, tt.input)
	}
}
