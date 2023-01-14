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
	Integer    = "Integer"

	// operators
	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"

	LessThan    = "<"
	GreaterThan = ">"

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
	If       = "If"
	Else     = "Else"
	Return   = "Return"
	True     = "True"
	False    = "False"
)
