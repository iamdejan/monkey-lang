package token

var keywords = map[string]TokenType{
	"fn":  Function,
	"let": Let,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
