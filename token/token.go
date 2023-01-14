package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "Illegal"
	Eof     = "Eof"

	// variable identifier and literal
	Identifier = "Identifier"
	Int        = "Int"

	// operators
	Assign = "="
	Plus   = "-"

	// delimiters
	Comma     = ","
	Semicolon = ";"

	LeftParenthesis  = "("
	RightParenthesis = ")"
	LeftBrace        = "{"
	RightBrace       = "}"

	// keywords
	Function = "Function"
	Let      = "Let"
)
