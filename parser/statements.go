package parser

func (p *Parser) ParseIf() *If {
	p.Advance()

	var cond Expression
	p.ParseExpression(&cond)

	return &If{nil, []Node{}, p.Line, p.Column, p.Filename, cond}
}

func (p *Parser) ParseWhile() *While {
	p.Advance()

	var cond Expression
	p.ParseExpression(&cond)

	return &While{nil, []Node{}, p.Line, p.Column, p.Filename, cond}
}

func (p *Parser) ParsePackage() *Package {
	p.Advance()
	nm := p.Current().Raw

	if p.Peek().Raw == "type" {
		// handle allocators and such
	}

	pkg := &Package{nil, nil, p.Line, p.Column, p.Filename, nm}
	return pkg
}
