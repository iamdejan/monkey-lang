package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

// region let statement

type LetTest struct {
	input              string
	expectedIdentifier string
	expectedValue      interface{}
}

func TestLetStatements(t *testing.T) {
	tests := []LetTest{
		{input: "let x = 5;", expectedIdentifier: "x", expectedValue: 5},
		{input: "let y = true;", expectedIdentifier: "y", expectedValue: true},
		{input: "let foobar = y;", expectedIdentifier: "foobar", expectedValue: "y"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong program.Statements length. expected=`1`, actual=`%d`", len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

// end region let statement

// region return statement

type ReturnTest struct {
	input         string
	expectedValue string
}

func TestReturnStatements(t *testing.T) {
	tests := []ReturnTest{
		{input: "return 5;", expectedValue: "5"},
		{input: "return 1 + 2;", expectedValue: "(1 + 2)"},
		{input: "return add(2, 3);", expectedValue: "add(2, 3)"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)

		program := p.ParseProgram()
		if program == nil {
			t.Fatal("`program` is null")
		}

		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong program.Statements length. expected=`1`, actual=`%d`", len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testReturnStatement(t, stmt, tt.expectedValue) {
			return
		}
	}
}

// end region return statement

// region identifier expression

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. expected=`1` statement, actual=`%d` statement(s).", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
	}

	testLiteralExpression(t, stmt.Expression, "foobar")
}

// end region identifier expression

// region integer literal

func TestIntegerLiteral(t *testing.T) {
	input := `5;`

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. expected=`1` statement, actual=`%d` statement(s).", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
	}

	testLiteralExpression(t, stmt.Expression, 5)
}

// end region integer literal

// region prefix expressions

type PrefixTest struct {
	input    string
	operator string
	value    interface{}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []PrefixTest{
		{input: "!5;", operator: "!", value: 5},
		{input: "-15;", operator: "-", value: 15},
		{input: "!true;", operator: "!", value: true},
		{input: "!false;", operator: "!", value: false},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong length for `program.Statements`. expected=`%d`, actual=`%d`", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("wrong type of `program.Statements[0]`. expected=`*ast.ExpressionStatement`, actual=`%T`", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("wrong type of `stmt.Expression`. expected=`*ast.PrefixExpression`, actual=`%T`", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("wrong `exp.Operator`. expected=`%s`, actual=`%s`", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

// end region prefix expressions

// region infix expressions

type InfixExpressionTest struct {
	input      string
	leftValue  interface{}
	operator   string
	rightValue interface{}
}

func TestParsingInfixExpressions(t *testing.T) {
	tests := []InfixExpressionTest{
		{input: "5 + 5", leftValue: 5, operator: "+", rightValue: 5},
		{input: "1 + 2", leftValue: 1, operator: "+", rightValue: 2},
		{input: "5 - 5", leftValue: 5, operator: "-", rightValue: 5},
		{input: "5 * 5", leftValue: 5, operator: "*", rightValue: 5},
		{input: "5 / 5", leftValue: 5, operator: "/", rightValue: 5},
		{input: "5 == 5", leftValue: 5, operator: "==", rightValue: 5},
		{input: "5 != 5", leftValue: 5, operator: "!=", rightValue: 5},
		{input: "5 < 5", leftValue: 5, operator: "<", rightValue: 5},
		{input: "5 > 5", leftValue: 5, operator: ">", rightValue: 5},
		{input: "5 <= 5", leftValue: 5, operator: "<=", rightValue: 5},
		{input: "5 >= 5", leftValue: 5, operator: ">=", rightValue: 5},
		{input: "a && b", leftValue: "a", operator: "&&", rightValue: "b"},
		{input: "c || d", leftValue: "c", operator: "||", rightValue: "d"},
		{input: "true == true", leftValue: true, operator: "==", rightValue: true},
		{input: "true != false", leftValue: true, operator: "!=", rightValue: false},
		{input: "false == false", leftValue: false, operator: "==", rightValue: false},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong program.Statements length. expected=`1`, actual=`%d`", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("wrong type for program.Statements[0]. expected=`*ast.ExpressionStatement`, actual=`%T`", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("wrong type for stmt.Expression. expected=`*ast.InfixExpression`, actual=`%T`", stmt.Expression)
		}

		testInfixExpression(t, exp, tt.leftValue, tt.operator, tt.rightValue)
	}
}

type OperatorPrecedenceTest struct {
	input    string
	expected string
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []OperatorPrecedenceTest{
		{input: "1 + 2 + 3", expected: "((1 + 2) + 3)"},
		{input: "1 * 2 + 3", expected: "((1 * 2) + 3)"},
		{input: "1 + 2 / 3", expected: "(1 + (2 / 3))"},
		{input: "-a * b", expected: "((-a) * b)"},
		{input: "!-a", expected: "(!(-a))"},
		{input: "a * b / c + d", expected: "(((a * b) / c) + d)"},
		{input: "3 + 4 * 5 == 3 * 1 + 4 * 5", expected: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{input: "true", expected: "true"},
		{input: "false", expected: "false"},
		{input: "3 > 5 == false", expected: "((3 > 5) == false)"},
		{input: "3 < 5 == true", expected: "((3 < 5) == true)"},
		{input: "1 + (2 + 3) + 4", expected: "((1 + (2 + 3)) + 4)"},
		{input: "(5 + 5) * 2", expected: "((5 + 5) * 2)"},
		{input: "2 / (5 + 5)", expected: "(2 / (5 + 5))"},
		{input: "-(5 + 5)", expected: "(-(5 + 5))"},
		{input: "!(true == true)", expected: "(!(true == true))"},
		{input: "a + add(b * c) + d", expected: "((a + add((b * c))) + d)"},
		{input: "add(a, b, 1, 2 * 3, 4 + 5)", expected: "add(a, b, 1, (2 * 3), (4 + 5))"},
		{input: "a * [1, 2, 3, 4][b * c] * d", expected: "((a * ([1, 2, 3, 4][(b * c)])) * d)"},
		{input: "add(a * b[2], b[1], 2 * [1, 2][1])", expected: "add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Fatalf("wrong evaluation. expected=`%s`, actual=`%s`", tt.expected, actual)
		}
	}
}

// end region infix expressions

// region boolean expressions

func TestBooleanLiteral(t *testing.T) {
	input := `true;`

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. expected=`1` statement, actual=`%d` statement(s).", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
	}

	testLiteralExpression(t, stmt.Expression, true)
}

// end region boolean expressions

// region if expressions

func TestIfExpression(t *testing.T) {
	input := "if (x < y) { x }"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. expected=`1` statement, actual=`%d` statement(s).", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("wrong type for stmt.Expression. exptected=`*ast.IfExpression`, actual=`%T`", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Fatalf("wrong length for exp.Consequence.Statements. expected=`1`, actual=`%d`", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("wrong type for exp.Consequence.Statements[0]. expected=`*ast.ExpressionStatement`, actual=`%T`", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative != nil {
		t.Fatalf("exp.Alternative should be null, but instead got `%#v`", exp.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := "if (x < y) { x } else { y }"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. expected=`1` statement, actual=`%d` statement(s).", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("wrong type for stmt.Expression. exptected=`*ast.IfExpression`, actual=`%T`", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Fatalf("wrong length for exp.Consequence.Statements. expected=`1`, actual=`%d`", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("wrong type for exp.Consequence.Statements[0]. expected=`*ast.ExpressionStatement`, actual=`%T`", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative == nil {
		t.Fatalf("exp.Alternative should not be null")
	}

	if len(exp.Alternative.Statements) != 1 {
		t.Fatalf("wrong length for exp.Alternative.Statements. expected=`1`, actual=`%d`", len(exp.Consequence.Statements))
	}

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("wrong type for exp.Alternative.Statements[0]. expected=`*ast.ExpressionStatement`, actual=`%T`", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}
}

// end region if expressions

// region function literal

func TestFunctionLiteralParsing(t *testing.T) {
	input := "fn(x, y) { x + y; }"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. expected=`1` statement, actual=`%d` statement(s).", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("wrong type for stmt.Expression. exptected=`*ast.FunctionLiteral`, actual=`%T`", stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("wrong function.Parameters length. expected=`2`, actual=`%d`", len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("wrong function.Body.Statements length. expected=`1`, actual=`%d`", len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("wrong type for function.Body.Statements[0]. expected=`*ast.ReturnStatement`, actual=`%T`", function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

type FunctionParamTest struct {
	input          string
	expectedParams []string
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []FunctionParamTest{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not `*ast.ExpressionStatement`, but rather `%T`", program.Statements[0])
		}

		function, ok := stmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("wrong type for stmt.Expression. exptected=`*ast.FunctionLiteral`, actual=`%T`", stmt.Expression)
		}

		if len(function.Parameters) != len(tt.expectedParams) {
			t.Fatalf("wrong function.Parameters length. expected=`%d`, actual=`%d`", len(tt.expectedParams), len(function.Parameters))
		}

		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

// end region function literal

// region call expression

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, c);"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("wrong program.Statements length. expected=`1`, actual=`%d`", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("wrong program.Statements[0] type. expected=`*ast.ExpressionStatement`, actual=`%T`", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("wrong stmt.Expression type. expected=`*ast.CallExpression`, actual=`%T`", stmt.Expression)
	}

	if !testIdentifier(t, exp.Function, "add") {
		return
	}

	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong exp.Arguments length. expected=`3`, actual=`%d`", len(exp.Arguments))
	}

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testIdentifier(t, exp.Arguments[2], "c")
}

// end region call expression

// region string literal expression

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("wrong type for `stmt`. expected=`*ast.StringLiteral`, actual=`%T`", stmt.Expression)
	}

	expected := "hello world"
	if literal.Value != expected {
		t.Fatalf("wrong `literal.Value`. expected=`%s`, actual=`%s`", expected, literal.Value)
	}
}

// end region string literal expression

// region array literal expression

func TestParsingArrayLiterals(t *testing.T) {
	input := `[1, 2 * 2, 3 + 3]`

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("wrong type for `stmt`. expected=`*ast.ArrayLiteral`, actual=`%T`", stmt.Expression)
	}

	if len(array.Elements) != 3 {
		t.Fatalf("wrong length for `array.Elements`. expected=`3`, actual=`%d`", len(array.Elements))
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

// end region array literal expression

func TestParsingIndexExpressions(t *testing.T) {
	input := "myArray[1 + 1]"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	stmt, _ := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("wrong type for `indexExp`. expected=`*ast.IndexExpression`, actual=`%T`", stmt.Expression)
	}

	if !testIdentifier(t, indexExp.Left, "myArray") {
		return
	}

	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}
}
