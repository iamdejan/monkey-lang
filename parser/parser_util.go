package parser

import (
	"fmt"
	"monkey/token"
)

func (p *Parser) peekError(expected token.TokenType) {
	msg := fmt.Sprintf("next token error. expected=`%s`, actual=`%s`", expected, p.peek.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
