package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readCharacter()
	return l
}

func (l *Lexer) readCharacter() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '!':
		if l.peekChar() == '=' {
			ch := l.character
			l.readCharacter()
			peek := l.character
			tok = token.Token{
				Type:    token.NotEqual,
				Literal: string(ch) + string(peek),
			}
		} else {
			tok = newToken(token.Bang, l.character)
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.character
			l.readCharacter()
			peek := l.character
			tok = token.Token{
				Type:    token.Equal,
				Literal: string(ch) + string(peek),
			}
		} else {
			tok = newToken(token.Assign, l.character)
		}
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
	case '-':
		tok = newToken(token.Minus, l.character)
	case '*':
		tok = newToken(token.Asterisk, l.character)
	case '/':
		tok = newToken(token.Slash, l.character)
	case '<':
		tok = newToken(token.LessThan, l.character)
	case '>':
		tok = newToken(token.GreaterThan, l.character)
	case '{':
		tok = newToken(token.LeftBrace, l.character)
	case '}':
		tok = newToken(token.RightBrace, l.character)
	case 0:
		tok = newToken(token.Eof, l.character)
	default:
		if isLetter(l.character) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.character) {
			tok.Literal = l.readDigit()
			tok.Type = token.Integer
			return tok
		} else {
			tok = newToken(token.Illegal, l.character)
		}
	}

	l.readCharacter()

	return tok
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.character) {
		l.readCharacter()
	}

	return l.input[startPos:l.position]
}

func (l *Lexer) readDigit() string {
	startPos := l.position
	for isDigit(l.character) {
		l.readCharacter()
	}

	return l.input[startPos:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}
