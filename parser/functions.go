package parser

import (
	"chroma/lexer"
	"chroma/logger"
	"strconv"

	"github.com/k0kubun/pp"
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

	return &Function{nil, []Node{}, p.Line, p.Column, p.Filename, name, tp, vars}
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

			ret = append(ret, Variable{nil, []Node{}, name, tp, p.Line, p.Column, p.Filename, nil /*THIS CAN NOT BE NIL FOREVER*/})
		case lexer.Comma:
			p.Advance()
		case lexer.RParen:
			reading = false
			p.Advance()
		default:
			logger.Exit("(Parsing) Failed to parse function arguments (Unrecognized Symbol)",
				[]string{"Line", strconv.Itoa(p.Line)},
				[]string{"Col", strconv.Itoa(p.Column)},
				[]string{"Symbol", p.Current().Raw},
				[]string{"Ast", "\n" + pp.Sprintln(ret) + "\n"},
				[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line, "(Parsing) Failed to parse function arguments (Unrecognized Symbol)") + "\n"})
		}

		if !reading {
			break
		}
	}

	return ret
}
