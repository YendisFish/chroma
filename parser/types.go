package parser

func (p *Parser) ReadType() TypeInfo {
	//attempt to read pointers and arrays first

	name := p.Current().Raw
	p.Advance()

	return TypeInfo{name, nil, 0, 0}
}
