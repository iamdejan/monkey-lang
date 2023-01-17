package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lexer   *lexer.Lexer
	current token.Token
	peek    token.Token
	errors  []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lexer.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.current.Type != token.Eof {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.current.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{
		Token: p.current,
	}

	if !p.expectPeek(token.Identifier) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.current,
		Value: p.current.Literal,
	}

	if !p.expectPeek(token.Assign) {
		return nil
	}

	// TODO: for now, skipping the expression(s) until we encounter semicolon
	for p.current.Type != token.Semicolon {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{
		Token: p.current,
	}

	p.nextToken()

	// TODO: for now, skipping the expression(s) until we encounter semicolon
	for p.current.Type != token.Semicolon {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) peekError(expected token.TokenType) {
	msg := fmt.Sprintf("next token error. expected=`%s`, actual=`%s`", expected, p.peek.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek.Type == t {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}
