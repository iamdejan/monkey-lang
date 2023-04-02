package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken_Operators(t *testing.T) {
	input := `=+(){},;`

	tests := []token.Token{
		{Type: token.Assign, Literal: "="},
		{Type: token.Plus, Literal: "+"},
		{Type: token.LeftParenthesis, Literal: "("},
		{Type: token.RightParenthesis, Literal: ")"},
		{Type: token.LeftBrace, Literal: "{"},
		{Type: token.RightBrace, Literal: "}"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Eof, Literal: ""},
	}

	testLexer(t, input, tests)
}

func TestNextToken_BasicCode(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x, y) {
	    x + y;
	}

	let result = add(five, ten);
	`

	tests := []token.Token{
		{Type: token.Let, Literal: "let"},
		{Type: token.Identifier, Literal: "five"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.Let, Literal: "let"},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.Let, Literal: "let"},
		{Type: token.Identifier, Literal: "add"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Function, Literal: "fn"},
		{Type: token.LeftParenthesis, Literal: "("},
		{Type: token.Identifier, Literal: "x"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Identifier, Literal: "y"},
		{Type: token.RightParenthesis, Literal: ")"},
		{Type: token.LeftBrace, Literal: "{"},
		{Type: token.Identifier, Literal: "x"},
		{Type: token.Plus, Literal: "+"},
		{Type: token.Identifier, Literal: "y"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.RightBrace, Literal: "}"},

		{Type: token.Let, Literal: "let"},
		{Type: token.Identifier, Literal: "result"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Identifier, Literal: "add"},
		{Type: token.LeftParenthesis, Literal: "("},
		{Type: token.Identifier, Literal: "five"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.RightParenthesis, Literal: ")"},
		{Type: token.Semicolon, Literal: ";"},
	}

	testLexer(t, input, tests)
}

func TestNextToken_ArithmeticAndLogicOperators(t *testing.T) {
	input := `
	!-*/5;
	5 < 10 > 5;
	`

	tests := []token.Token{
		{Type: token.Bang, Literal: "!"},
		{Type: token.Minus, Literal: "-"},
		{Type: token.Asterisk, Literal: "*"},
		{Type: token.Slash, Literal: "/"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.Integer, Literal: "5"},
		{Type: token.LessThan, Literal: "<"},
		{Type: token.Integer, Literal: "10"},
		{Type: token.GreaterThan, Literal: ">"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
	}

	testLexer(t, input, tests)
}

func TestNextToken_BranchingAndReturnKeywords(t *testing.T) {
	input := `
	if (5 < 10) {
	    return true;
	} else {
	    return false;
	}
	`

	tests := []token.Token{
		{Type: token.If, Literal: "if"},
		{Type: token.LeftParenthesis, Literal: "("},
		{Type: token.Integer, Literal: "5"},
		{Type: token.LessThan, Literal: "<"},
		{Type: token.Integer, Literal: "10"},
		{Type: token.RightParenthesis, Literal: ")"},
		{Type: token.LeftBrace, Literal: "{"},
		{Type: token.Return, Literal: "return"},
		{Type: token.True, Literal: "true"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.RightBrace, Literal: "}"},
		{Type: token.Else, Literal: "else"},
		{Type: token.LeftBrace, Literal: "{"},
		{Type: token.Return, Literal: "return"},
		{Type: token.False, Literal: "false"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.RightBrace, Literal: "}"},
	}

	testLexer(t, input, tests)
}

func TestNextToken_BooleanOperators(t *testing.T) {
	input := `
	10 == 10;
	ten != 30;
	`

	tests := []token.Token{
		{Type: token.Integer, Literal: "10"},
		{Type: token.Equal, Literal: "=="},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.NotEqual, Literal: "!="},
		{Type: token.Integer, Literal: "30"},
		{Type: token.Semicolon, Literal: ";"},
	}

	testLexer(t, input, tests)
}

func testLexer(t *testing.T, input string, tests []token.Token) {
	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - wrong token type. expected = %q, got = %q -> lexer = %#v", i, tt.Type, tok.Type, l)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - wrong literal. expected = %q, got = %q -> lexer = %#v", i, tt.Literal, tok.Literal, l)
		}
	}
}

func TestNextToken_String(t *testing.T) {
	input := `
	"foobar"
	"foo bar"
	"12345"
	`

	tests := []token.Token{
		{Type: token.String, Literal: "foobar"},
		{Type: token.String, Literal: "foo bar"},
		{Type: token.String, Literal: "12345"},
	}

	testLexer(t, input, tests)
}

func TestNextToken_Brackets(t *testing.T) {
	input := `[1, 2]`;

	tests := []token.Token{
		{Type: token.LeftBracket, Literal: "["},
		{Type: token.Integer, Literal: "1"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Integer, Literal: "2"},
		{Type: token.RightBracket, Literal: "]"},
	}

	testLexer(t, input, tests)
}
