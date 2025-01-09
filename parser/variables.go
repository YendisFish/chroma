package parser

import (
	"chroma/lexer"
)

func (p *Parser) ParseVariable() *Variable {
	p.Advance()
	name := p.Current().Raw
	p.Advance()

	var tp TypeInfo
	if p.Current().Type != lexer.Equals {
		tp = p.ReadType()
	}

	var expr Expression
	if p.Current().Type == lexer.Equals {
		p.Advance()
		p.ParseExpression(&expr)
	}

	return &Variable{nil, []Node{}, name, tp, p.Line, p.Column, p.Filename, expr}
}
