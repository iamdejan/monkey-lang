package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

const (
	_ int = iota
	Lowest
	Equals        // e.g. 1 == a
	LessOrGreater // e.g. 2 < 3 or 3 > 1
	Sum           // e.g. 2 + 4
	Product       // e.g. 5 * 3
	Prefix        // e.g. --5
	Call          // e.g. add(2, 3)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression // left side is the argument
)

type Parser struct {
	lexer  *lexer.Lexer
	errors []string

	current token.Token
	peek    token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.Identifier, p.parseIdentifier)
	p.registerPrefix(token.Integer, p.parseIntegerLiteral)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.current,
		Value: p.current.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{
		Token: p.current,
	}

	value, err := strconv.ParseInt(p.current.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.current.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	literal.Value = value

	return literal
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lexer.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
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
		return p.parseExpressionStatement()
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.current,
	}

	stmt.Expression = p.parseExpression(Lowest)

	if p.peek.Type == token.Semicolon {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.current.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
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
