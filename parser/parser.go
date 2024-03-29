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
	Equal         // e.g. 1 == a
	LessOrGreater // e.g. 2 < 3 or 3 > 1
	Boolean       // e.g. a && b or c || d
	Sum           // e.g. 2 + 4
	Product       // e.g. 5 * 3
	Prefix        // e.g. -5
	Call          // e.g. add(2, 3)
	Index         // e.g. myArray[5]
)

var precedences = map[token.TokenType]int{
	token.Equal:              Equal,
	token.NotEqual:           Equal,
	token.LessThan:           LessOrGreater,
	token.GreaterThan:        LessOrGreater,
	token.LessThanOrEqual:    LessOrGreater,
	token.GreaterThanOrEqual: LessOrGreater,
	token.BooleanAnd:         Boolean,
	token.BooleanOr:          Boolean,
	token.Plus:               Sum,
	token.Minus:              Sum,
	token.Slash:              Product,
	token.Asterisk:           Product,
	token.LeftParenthesis:    Call,
	token.LeftBracket:        Index,
}

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
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.True, p.parseBoolean)
	p.registerPrefix(token.False, p.parseBoolean)
	p.registerPrefix(token.LeftParenthesis, p.parseGroupedExpression)
	p.registerPrefix(token.If, p.parseIfExpression)
	p.registerPrefix(token.Function, p.parseFunctionLiteral)
	p.registerPrefix(token.String, p.parseStringLiteral)
	p.registerPrefix(token.LeftBracket, p.parseArrayLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Slash, p.parseInfixExpression)
	p.registerInfix(token.Asterisk, p.parseInfixExpression)
	p.registerInfix(token.Equal, p.parseInfixExpression)
	p.registerInfix(token.NotEqual, p.parseInfixExpression)
	p.registerInfix(token.LessThan, p.parseInfixExpression)
	p.registerInfix(token.LessThanOrEqual, p.parseInfixExpression)
	p.registerInfix(token.GreaterThan, p.parseInfixExpression)
	p.registerInfix(token.GreaterThanOrEqual, p.parseInfixExpression)
	p.registerInfix(token.BooleanAnd, p.parseInfixExpression)
	p.registerInfix(token.BooleanOr, p.parseInfixExpression)
	p.registerInfix(token.LeftParenthesis, p.parseCallExpression)
	p.registerInfix(token.LeftBracket, p.parseIndexExpression)

	// read 2 tokens, so that current and peek (next) token are set
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

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.current,
		Operator: p.current.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(Prefix)

	return expression
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

	p.nextToken()

	stmt.Value = p.parseExpression(Lowest)

	if p.peek.Type == token.Semicolon {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{
		Token: p.current,
	}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(Lowest)
	if p.peek.Type == token.Semicolon {
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
		p.noPrefixParseFnError(p.current.Type)
		return nil
	}

	leftExp := prefix()

	for p.peek.Type != token.Semicolon && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peek.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek.Type == t {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peek.Type]; ok {
		return p
	}

	return Lowest
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.current.Type]; ok {
		return p
	}

	return Lowest
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.current,
		Left:     left,
		Operator: p.current.Literal,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.BooleanLiteral{
		Token: p.current,
		Value: (p.current.Type == token.True),
	}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(Lowest)

	if !p.expectPeek(token.RightParenthesis) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	exp := &ast.IfExpression{
		Token: p.current,
	}

	if !p.expectPeek(token.LeftParenthesis) {
		return nil
	}

	p.nextToken()

	exp.Condition = p.parseExpression(Lowest)

	if !p.expectPeek(token.RightParenthesis) {
		return nil
	}

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	exp.Consequence = p.parseBlockStatement()

	// `else` (alternative) block handling
	if p.peek.Type == token.Else {
		p.nextToken()

		if !p.expectPeek(token.LeftBrace) {
			return nil
		}

		exp.Alternative = p.parseBlockStatement()
	}

	return exp
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token: p.current,
	}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for p.current.Type != token.RightBrace && p.current.Type != token.Eof {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{
		Token: p.current,
	}

	if !p.expectPeek(token.LeftParenthesis) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peek.Type == token.RightParenthesis {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{
		Token: p.current,
		Value: p.current.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peek.Type == token.Comma {
		p.nextToken()
		p.nextToken()

		ident := &ast.Identifier{
			Token: p.current,
			Value: p.current.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RightParenthesis) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{
		Token:    p.current,
		Function: function,
	}
	exp.Arguments = p.parseExpressionList(token.RightParenthesis)
	return exp
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.current,
		Value: p.current.Literal,
	}
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	al := &ast.ArrayLiteral{
		Token: p.current,
	}

	al.Elements = p.parseExpressionList(token.RightBracket)

	return al
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	expressions := []ast.Expression{}

	if p.peek.Type == end {
		p.nextToken()
		return expressions
	}

	p.nextToken()
	expressions = append(expressions, p.parseExpression(Lowest))
	for p.peek.Type == token.Comma {
		p.nextToken()
		p.nextToken()

		expressions = append(expressions, p.parseExpression(Lowest))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return expressions
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{
		Token: p.current,
		Left:  left,
	}

	p.nextToken()
	exp.Index = p.parseExpression(Lowest)

	if !p.expectPeek(token.RightBracket) {
		return nil
	}

	return exp
}
