package parser

import (
	"chroma/lexer"
)

func (p *Parser) ParseFunction() *Function {
	p.Advance()

	var argOne *Variable = nil
	ifunc := false
	if p.Current().Type == lexer.LParen {
		p.Advance()

		instanceName := p.Current().Raw
		p.Advance()

		instanceTp := p.ReadType()

		argOne = &Variable{nil, []Node{}, instanceName, instanceTp, p.Line, p.Column, p.Filename, nil}
		ifunc = true

		p.Advance()
	}

	var name string
	special := false
	switch tp := p.Current(); tp.Type {
	case lexer.LBrack:
		//handle indexer
		special = true
		name = "__internal_index_"
		p.Advance()
	case lexer.Word:
		if tp.Raw == "delete" {
			//handle free function
			break
		}

		if tp.Raw == "alloc" {
			//handle alloc function
			break
		}

		name = p.Current().Raw
	}

	p.Advance()
	var vars []Variable
	vars = p.ParseFunctionArgs()

	if ifunc {
		vars = append([]Variable{*argOne}, vars...)
	}

	var tp TypeInfo
	if p.Current().Type != lexer.LBrace {
		tp = p.ReadType()

		if special {
			name = name + "get"
		}
	} else {
		tp = TypeInfo{nil, nil, nil}

		if special {
			name = name + "set"
		}
	}

	return &Function{nil, []Node{}, p.Line, p.Column, p.Filename, name, tp, vars, ifunc}
}

func (p *Parser) ParseFunctionArgs() []Variable {
	//start on (
	p.Advance()

	ret := []Variable{}
	for reading := true; reading; {
		switch tok := p.Current(); tok.Type {
		case lexer.Word:
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
			p.Panic("Unrecognized symbol", "Function arguments")
		}

		if !reading {
			break
		}
	}

	return ret
}
