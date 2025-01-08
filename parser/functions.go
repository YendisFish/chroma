package parser

import (
	"chroma/lexer"
)

func (p *Parser) ParseFunction() *Function {
	p.Advance()
	name := p.Current().Raw
	p.Advance()
	vars := p.ParseFunctionArgs()

	var tp TypeInfo
	if p.Current().Type != lexer.LBrace {
		tp = p.ReadType()
	} else {
		tp = TypeInfo{}
	}

	return &Function{nil, []Node{}, name, tp, vars}
}

func (p *Parser) ParseFunctionArgs() []Variable {
	//start on (
	p.Advance()

	ret := []Variable{}
	for reading := true; reading; {
		switch tok := p.Current(); tok.Type {
		case lexer.Word:
			if p.Peek().Type == lexer.Colon {
				//handle var declaration
			}

			name := p.Current().Raw
			p.Advance()
			tp := p.ReadType()

			ret = append(ret, Variable{nil, []Node{}, name, tp})
		case lexer.Comma:
			p.Advance()
		case lexer.RParen:
			reading = false
			p.Advance()
		}

		if !reading {
			break
		}
	}

	return ret
}
