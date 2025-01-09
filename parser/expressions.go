package parser

import (
	"chroma/lexer"
	"chroma/logger"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func (p *Parser) ParseExpression(expr *Expression) {
	for reading := true; reading; {
		switch p.Current().Type {
		case lexer.Word:
			if val, ok := determineLiteralType(p.Current()); ok {
				*expr = &Literal{nil, nil, p.Line, p.Column, p.Filename, p.Current().Raw, val}
				p.Advance()
			}

			if p.Current().Type == lexer.Word {
				p.Advance()
				reading = false
			}
		case lexer.Add, lexer.Subtract, lexer.Star, lexer.BSlash:
			//check for boolop
			//parse binop
		case lexer.LBrace:
			reading = false
		default:
			logger.Exit("(Parsing) Unrecognized token",
				[]string{"Line", strconv.Itoa(p.Line)},
				[]string{"Col", strconv.Itoa(p.Column)},
				[]string{"Token", p.Current().Raw},
				[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
				[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line, "(Parsing) Unrecognized token") + "\n"})
		}

		if !reading {
			break
		}
	}
}

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
