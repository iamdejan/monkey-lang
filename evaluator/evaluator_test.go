package evaluator

import "testing"

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
