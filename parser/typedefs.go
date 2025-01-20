package parser

import (
	"chroma/lexer"
)

func (p *Parser) ParseTypeDef() *TypeDef {
	p.Advance()

	nm := p.Current().Raw

	p.Advance()

	var tp TypeType
	var pnt *string = nil
	switch p.Current().Raw {
	case "struct":
		tp = TStruct
	case "interface":
		tp = TInterface
	default:
		tp = TInherit

		pnt = new(string)
		*pnt = p.Current().Raw
	}

	p.Advance()

	var ret *TypeDef = &TypeDef{nil, []Node{}, p.Line, p.Column, p.Filename, nm, tp, pnt}

	if p.Current().Type == lexer.LBrace {
		p.Advance()

		for rding := true; rding; {
			switch p.Current().Type {
			case lexer.Word:
				nm := p.Current().Raw
				p.Advance()
				tp := p.ReadType()

				ret.children = append(ret.children, &Variable{ret, nil, nm, tp, p.Line, p.Column, p.Filename, nil})
			case lexer.RBrace:
				rding = false
			default:
				p.Panic("Unrecognized symbol", "Struct Parsing")
			}

			if !rding {
				break
			}
		}
	}

	return ret
}
