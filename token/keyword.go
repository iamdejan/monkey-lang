package token

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"if":     If,
	"else":   Else,
	"return": Return,
	"true":   True,
	"false":  False,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
