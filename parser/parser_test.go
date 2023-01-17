package parser

import (
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
		t.Fatalf("statement is not `*ast.LetStatement`, but rather `%T`", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Fatalf("incorrect variable name. expected=`%s`, actual=`%s`", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Fatalf("incorrect letStmt.Name.TokenLiteral(). expected=`%s`, actual=`%s`", name, letStmt.Name.TokenLiteral())
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
