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

	return &Variable{nil, []Node{}, name, tp}
}
