package lexer

import (
	"testing"

	"monkey/token"
)

type TokenTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken_Operators(t *testing.T) {
	input := `=+(){},;`

	tests := []TokenTest{
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.LeftParenthesis, "("},
		{token.RightParenthesis, ")"},
		{token.LeftBrace, "{"},
		{token.RightBrace, "}"},
		{token.Comma, ","},
		{token.Semicolon, ";"},
		{token.Eof, ""},
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

	tests := []TokenTest{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},

		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},

		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LeftParenthesis, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RightParenthesis, ")"},
		{token.LeftBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},

		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.LeftParenthesis, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RightParenthesis, ")"},
		{token.Semicolon, ";"},
	}

	testLexer(t, input, tests)
}

func TestNextToken_ArithmeticAndLogicOperators(t *testing.T) {
	input := `
	!-*/5;
	5 < 10 > 5;
	`

	tests := []TokenTest{
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Asterisk, "*"},
		{token.Slash, "/"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},

		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.GreaterThan, ">"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
	}

	testLexer(t, input, tests)
}

func testLexer(t *testing.T, input string, tests []TokenTest) {
	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected = %q, got = %q -> lexer = %#v", i, tt, tok, l)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected = %q, got = %q -> lexer = %#v", i, tt, tok, l)
		}
	}
}
