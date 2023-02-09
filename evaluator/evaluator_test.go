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
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
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
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
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
		testBooleanObject(t, evaluated, tt.expected)
	}
}
