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

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_SubsetCode() {t *testing.T} {
	// TODO dejan
}
