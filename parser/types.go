package parser

import (
	"chroma/lexer"
	"chroma/logger"
	"strconv"

	"github.com/k0kubun/pp"
)

func (p *Parser) ReadType() TypeInfo {
	var ptrs []Ptr = nil
	if p.Current().Type == lexer.Star || p.Current().Type == lexer.LBrack {
		ptrs = []Ptr{}
		for reading := true; reading; {
			switch p.Current().Type {
			case lexer.Star:
				ptrs = append(ptrs, Ptr{Pointer, nil})
				p.Advance()
			case lexer.LBrack:
				if p.Peek().Type == lexer.Word {
					p.Advance()
					index, err := strconv.ParseInt(p.Current().Raw, 0, 32)
					if err != nil {
						logger.Exit("(Parsing) Could not convert symbol to integer",
							[]string{"Line", strconv.Itoa(p.Line)},
							[]string{"Col", strconv.Itoa(p.Column)},
							[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
							[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line) + "\n"})
					}

					ptrs = append(ptrs, Ptr{Array, &index})

					p.Advance()
					p.Advance()
				} else {
					ptrs = append(ptrs, Ptr{Slice, nil})

					p.Advance()
					p.Advance()
				}
			case lexer.Word:
				reading = false
			default:
				logger.Exit("(Parsing) Failed to parse variable type (Unrecognized Symbol)",
					[]string{"Line", strconv.Itoa(p.Line)},
					[]string{"Col", strconv.Itoa(p.Column)},
					[]string{"Symbol", p.Current().Raw},
					[]string{"Ast", "\n" + pp.Sprintln(ptrs) + "\n"},
					[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line) + "\n"})
			}

			if !reading {
				break
			}
		}
	}

	name := p.Current().Raw
	p.Advance()

	var generics []TypeInfo = nil

	if p.Current().Type == lexer.LBrack {
		p.Advance()

		for reading := true; reading; {
			switch p.Current().Type {
			case lexer.Word:
				tp := p.ReadType()
				generics = append(generics, tp)
			}
		}
	}

	return TypeInfo{&name, generics, ptrs}
}
