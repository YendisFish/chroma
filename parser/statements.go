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

func (p *Parser) ParseImport() *Import {
	p.Advance()
	nm := p.Current().Raw

	return &Import{nil, nil, p.Line, p.Column, p.Filename, nm}
}

func (p *Parser) ParseForRange() *For {
	asn := &Literal{nil, nil, p.Line, p.Column, p.Filename, "0", LNum}
	tpnm := "i32"
	index := &Variable{p.node, nil, p.Current().Raw, TypeInfo{&tpnm, nil, nil}, p.Line, p.Column, p.Filename, asn}

	p.AdvanceBy(2)

	//see if we need to instanciate v later
	var v *Variable = nil
	vis := false
	nm := p.Current().Raw
	if nm != "_" {
		vis = true
	}

	p.AdvanceBy(2)

	//we read the expr after `range`
	var expr Expression
	p.ParseExpression(&expr)

	//generate a condition to handle the bounds of expr
	var cond Expression = nil //generate a condition one day
	ret := &For{nil, []Node{}, p.Line, p.Column, p.Filename, index, cond, nil /*one day, one day, this will be a pseudo operation*/}

	//if we have a value, we assign it to expr[i]
	if vis {
		switch tp := expr.(type) {
		case *FuncRef:
			tp.IndexRef = []Expression{
				&FuncRef{
					nil,
					nil,
					p.Line,
					p.Column,
					p.Filename,
					index.Name,
					nil,
				},
			}
		case *VarRef:
			tp.IndexRef = []Expression{
				&VarRef{
					nil,
					nil,
					p.Line,
					p.Column,
					p.Filename,
					index.Name,
					nil,
				},
			}
		}

		v = &Variable{ret, nil, nm, TypeInfo{nil, nil, nil}, p.Line, p.Column, p.Filename, expr}
		v.parent = ret
		ret.children = append(ret.children, v)
	}

	return ret
}
