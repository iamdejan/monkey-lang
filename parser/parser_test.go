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

func TestLetStatements_InvalidIdentifiers(t *testing.T) {
	input := `
	let x 5;
	let 1+2;
	let = add(3, 4);
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("`program` is null")
	}

	expectedErrors := []string{
		"next token error. expected=`=`, actual=`Integer`",
		"next token error. expected=`Identifier`, actual=`Integer`",
		"next token error. expected=`Identifier`, actual=`=`",
	}
	validateParseErrors(t, p, expectedErrors)
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

// region utilities

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
}

func validateParseErrors(t *testing.T, p *Parser, expectedErrors []string) {
	actualErrors := p.Errors()
	length := len(actualErrors)
	if length != len(expectedErrors) {
		t.Fatalf("invalid parser errors. expected=`%d` errors, actual=`%d` errors", len(expectedErrors), length)
		return
	}

	for i, err := range expectedErrors {
		if actualErrors[i] != err {
			t.Fatalf("invalid error message at %d. expected=`%s`, actual=`%s`", i, err, actualErrors[i])
			return
		}
	}
}

// end region utilities

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

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression is not `*ast.Identifier`, but rather `%T`", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("wrong ident.Value. expected=`foobar`, actual=`%s`", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("wrong ident.TokenLiteral(). expected=`foobar`, actual=`%s`", ident.TokenLiteral())
	}
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

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not `*ast.IntegerLiteral`, but rather `%T`", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("wrong literal.Value. expected=`5`, actual=`%d`", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("wrong literal.TokenLiteral(). expected=`5`, actual=`%s`", literal.TokenLiteral())
	}
}

// end region integer literal

// region prefix expressions

type PrefixTest struct {
	input        string
	operator     string
	integerValue int64
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []PrefixTest{
		{input: "!5;", operator: "!", integerValue: 5},
		{input: "-15;", operator: "-", integerValue: 15},
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

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
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

// end region prefix expressions
