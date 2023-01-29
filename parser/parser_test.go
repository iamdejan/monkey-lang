package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

// region let statement

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 1+2;
	let foobar = add(3, 4);
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("`program` is null")
	}

	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatal("`program` should have 3 statements")
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !correctLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func correctLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("incorrect token literal. expected=`let`, actual=`%s`", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not `*ast.LetStatement`, but rather `%T`", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("incorrect variable name. expected=`%s`, actual=`%s`", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("incorrect letStmt.Name.TokenLiteral(). expected=`%s`, actual=`%s`", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

// end region let statement

// region return statement

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 1+2;
	return add(2,3);
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("`program` is null")
	}

	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatal("`program` should have 3 statements")
	}

	for i, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("[%d] wrong stmt type. expected=`*ast.ReturnStatement`, actual=`%T`", i, stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("wrong returnStmt.TokenLiteral(). expected=`return`, actual=`%s`", returnStmt.TokenLiteral())
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

func TestOperatorPrecedence(t *testing.T) {
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

// region test helper functions

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	length := len(errors)
	if length == 0 {
		return
	}

	t.Errorf("parser has %d errors", length)
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case bool:
		return testBooleanLiteral(t, exp, bool(v))
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}
	t.Fatalf("type of exp not handled. actual=`%T`", exp)
	return false
}

func testBooleanLiteral(t *testing.T, bl ast.Expression, value bool) bool {
	b, ok := bl.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("wrong boolean literal type. expected=`*ast.BooleanLiteral`, got=`%T`", bl)
		return false
	}

	if b.Value != value {
		t.Errorf("wrong `b.Value`. expected=`%v`, got=`%v`", value, b.Value)
		return false
	}

	if b.TokenLiteral() != fmt.Sprintf("%v", value) {
		t.Errorf("wrong `b.TokenLiteral()`. expected=`%s`, got=`%s`", fmt.Sprintf("%v", value), b.TokenLiteral())
		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("wrong integer literal type. expected=`*ast.IntegerLiteral`, got=`%T`", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("wrong `integ.Value`. expected=`%d`, got=`%d`", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("wrong `integ.TokenLiteral()`. expected=`%s`, got=`%s`", fmt.Sprintf("%d", value), integ.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("wrong exp type. expected=`*ast.Identifier`, actual=`%T`", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("wrong ident.Value. expected=`%s`, actual=`%s`", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("wrong ident.TokenLiteral(). expected=`%s`, actual=%s", value, ident.TokenLiteral())
	}

	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("wrong exp type. expected=`*ast.InfixExpression`, actual=`%T(%s)`", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("wrong opExp.Operator. expected=`%s`, actual=`%s`", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

// end region test helper functions
