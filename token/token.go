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

	LessThan           = "<"
	GreaterThan        = ">"
	Equal              = "=="
	NotEqual           = "!="
	LessThanOrEqual    = "<="
	GreaterThanOrEqual = ">="
	BooleanAnd         = "&&"
	BooleanOr          = "||"

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

	String = "String"

	// array
	LeftBracket  = "["
	RightBracket = "]"
)
