package parser

import (
	"chroma/lexer"
	"strconv"
	"strings"
)

// parses an expression
func (p *Parser) ParseExpression(expr *Expression) {
	for reading := true; reading; {
		switch tok := p.Current(); tok.Type {
		case lexer.Word:
			if p.Peek().Type == lexer.Colon {
				vr := p.ParseShortVar()
				expr = nil
				reading = false

				p.Append(vr)
				break
			}

			switch tok.Raw {
			//ideally there will be cases for from, error, and other stuff as well
			case "from":
				p.Advance()
				var ex Expression
				p.ParseExpression(&ex)
				// parse the from expression

				var indexes []Expression = nil
				if p.Current().Type == lexer.LBrack {
					p.Advance()
					indexes = p.ParseIndexAccess()
				}

				*expr = &From{nil, []Node{}, p.Line, p.Column, p.Filename, ex, indexes}
			default:
				if val, ok := determineLiteralType(p.Current()); ok {
					*expr = &Literal{nil, nil, p.Line, p.Column, p.Filename, p.Current().Raw, val}
					p.Advance()
				} else if val, ok := p.determineFuncRef(); ok {
					*expr = val
					p.Advance()
				} else {
					val := p.parseVarRef()
					*expr = val
					p.Advance()
				}
			}

			// if p.Current().Line != tok.Line {
			// 	reading = false
			// }
			//p.Advance()
			if p.ExpressionEnd() {
				reading = false
			}
		case lexer.Add, lexer.Subtract, lexer.Star, lexer.FSlash:
			//check for boolop
			//parse binop

			binop := &BinOp{nil, nil, p.Line, p.Column, p.Filename, *expr, nil, p.Current().Raw}
			p.Advance()

			var right Expression
			p.ParseExpression(&right)
			binop.Right = right

			//p.Advance()

			if p.ExpressionEnd() {
				reading = false
			}

			*expr = binop
		case lexer.LParen:
			p.Advance()
		case lexer.RParen:
			p.Advance()
		case lexer.LBrack:
			//create index access expression eventually
			reading = false
		case lexer.LBrace, lexer.RBrack:
			//p.Advance()
			reading = false // for some reason this does what I want it to do... but I actually have no clue why!!

			//explanation
			//this function will skip the correct amount of paranthesis IE ((()))
			//and I have not told it to do so... idk where in the function it does it but it does
		default:
			p.Panic("Unrecognized token", "Expression")
		}

		if !reading {
			break
		}
	}
}

// parse, compress, and determine a literal
func determineLiteralType(tok lexer.Token) (LiteralType, bool) {
	if strings.HasPrefix(tok.Raw, "0x") {
		return LByte, true
	}

	if tok.Raw == "true" {
		return LTrue, true
	}

	if tok.Raw == "false" {
		return LFalse, true
	}

	_, err := strconv.ParseInt(tok.Raw, 0, 64)
	if err == nil {
		return LNum, true
	}

	return -1, false
}

// determine if reading a func and parse if reading
func (p *Parser) determineFuncRef() (*FuncRef, bool) {
	if pkt := p.Peek().Type; pkt != lexer.LBrack || pkt != lexer.LParen {
		return nil, false
	}

	nm := p.Current().Raw
	return &FuncRef{nil, nil, p.Line, p.Column, p.Filename, nm, nil}, true
}

// parse a variable
func (p *Parser) parseVarRef() *VarRef {
	nm := p.Current().Raw
	return &VarRef{nil, nil, p.Line, p.Column, p.Filename, nm, nil}
}

func (p *Parser) ParseShortVar() *Variable {
	nm := p.Current().Raw

	p.AdvanceBy(3)
	var expr Expression
	p.ParseExpression(&expr)

	return &Variable{p.node, []Node{}, nm, TypeInfo{nil, nil, nil}, p.Line, p.Column, p.Filename, expr}
}

func (p *Parser) ParseIndexAccess() []Expression {
	var ret []Expression = []Expression{}
	for r := true; r; {
		switch p.Current().Type {
		case lexer.RBrack:
			p.Advance()
			r = false
		default:
			var expr Expression
			p.ParseExpression(&expr)

			ret = append(ret, expr)
		}

		if !r {
			break
		}
	}

	return ret
}

func (p *Parser) ExpressionEnd() bool {
	switch p.Current().Type {
	case lexer.Add, lexer.Subtract, lexer.Star, lexer.FSlash, lexer.RParen, lexer.RBrack:
		return false
	default:
		return true
	}
}
