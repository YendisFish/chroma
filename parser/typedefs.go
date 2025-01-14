package parser

func (p *Parser) ParseTypeDef() *TypeDef {
	p.Advance()

	nm := p.Current().Raw

	p.Advance()

	var tp TypeType
	switch p.Current().Raw {
	case "struct":
		tp = TStruct
	case "interface":
		tp = TInterface
	default:
		tp = TInherit
	}

	p.Advance()

	return &TypeDef{nil, []Node{}, p.Line, p.Column, p.Filename, nm, tp}
}
