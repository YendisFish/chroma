package parser

import (
	"chroma/lexer"
	"strconv"
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
						p.Panic("Could not convert symbol into integer", "Variable Types")
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
				p.Panic("Unrecognized symbol", "Variable Types")
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
