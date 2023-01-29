package parser

import (
	"fmt"
	"monkey/ast"
	"testing"
)

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	length := len(errors)
	if length == 0 {
		return
	}

	t.Errorf("parser has %d errors", length)
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case bool:
		return testBooleanLiteral(t, exp, bool(v))
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}
	t.Fatalf("type of exp not handled. actual=`%T`", exp)
	return false
}

func testBooleanLiteral(t *testing.T, bl ast.Expression, value bool) bool {
	b, ok := bl.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("wrong boolean literal type. expected=`*ast.BooleanLiteral`, got=`%T`", bl)
		return false
	}

	if b.Value != value {
		t.Errorf("wrong `b.Value`. expected=`%v`, got=`%v`", value, b.Value)
		return false
	}

	if b.TokenLiteral() != fmt.Sprintf("%v", value) {
		t.Errorf("wrong `b.TokenLiteral()`. expected=`%s`, got=`%s`", fmt.Sprintf("%v", value), b.TokenLiteral())
		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("wrong integer literal type. expected=`*ast.IntegerLiteral`, got=`%T`", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("wrong `integ.Value`. expected=`%d`, got=`%d`", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("wrong `integ.TokenLiteral()`. expected=`%s`, got=`%s`", fmt.Sprintf("%d", value), integ.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("wrong exp type. expected=`*ast.Identifier`, actual=`%T`", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("wrong ident.Value. expected=`%s`, actual=`%s`", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("wrong ident.TokenLiteral(). expected=`%s`, actual=%s", value, ident.TokenLiteral())
	}

	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("wrong exp type. expected=`*ast.InfixExpression`, actual=`%T(%s)`", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("wrong opExp.Operator. expected=`%s`, actual=`%s`", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}
