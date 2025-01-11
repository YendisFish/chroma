package parser

import (
	"chroma/lexer"
	"chroma/logger"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

// parses an expression
func (p *Parser) ParseExpression(expr *Expression) {
	for reading := true; reading; {
		switch tok := p.Current(); tok.Type {
		case lexer.Word:
			switch tok.Raw {
			//ideally there will be cases for from, error, and other stuff as well
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

			if p.Current().Line != tok.Line {
				reading = false
			}

			//p.Advance()
		case lexer.Add, lexer.Subtract, lexer.Star, lexer.BSlash:
			//check for boolop
			//parse binop

			bop := &BinOp{nil, nil, p.Line, p.Column, p.Filename, *expr, nil, p.Current().Raw}
			p.Advance()

			var rght Expression
			p.ParseExpression(&rght)
			bop.Right = rght

			p.Advance()

			*expr = bop
		case lexer.LParen:
			p.Advance()
		case lexer.LBrace, lexer.RParen:
			//p.Advance()
			reading = false // for some reason this does what I want it to do... but I actually have no clue why!!

			//explanation
			//this function will skip the correct amount of paranthesis IE ((()))
			//and I have not told it to do so... idk where in the function it does it but it does
		default:
			logger.Exit("(Parsing)(EXPR) Unrecognized token",
				[]string{"Line", strconv.Itoa(p.Line)},
				[]string{"Col", strconv.Itoa(p.Column)},
				[]string{"Token", p.Current().Raw},
				[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
				[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line, "(Parsing)(EXPR) Unrecognized token") + "\n"})
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
