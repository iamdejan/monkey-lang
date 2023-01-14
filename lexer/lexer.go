package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readCharacter()
	return l
}

func (l *Lexer) readCharacter() {
	if l.position >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.position]
	}

	l.position += 1
	l.readPosition = l.position + 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.character {
	case '=':
		tok = newToken(token.Assign, l.character)
	case ';':
		tok = newToken(token.Semicolon, l.character)
	case '(':
		tok = newToken(token.LeftParenthesis, l.character)
	case ')':
		tok = newToken(token.RightParenthesis, l.character)
	case ',':
		tok = newToken(token.Comma, l.character)
	case '+':
		tok = newToken(token.Plus, l.character)
	case '{':
		tok = newToken(token.LeftBrace, l.character)
	case '}':
		tok = newToken(token.RightBrace, l.character)
	case 0:
		tok = newToken(token.Eof, l.character)
	}

	l.readCharacter()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: lit(ch),
	}
}

func lit(ch byte) string {
	if ch == 0 {
		return ""
	}

	return string(ch)
}
