package lexer

var symbols = map[rune]Token{
	'+':  Token{Type: Add},
	'-':  Token{Type: Subtract},
	'/':  Token{Type: FSlash},
	'*':  Token{Type: Star},
	'\\': Token{Type: BSlash},
	'{':  Token{Type: LBrace},
	'}':  Token{Type: RBrace},
	'[':  Token{Type: LBrack},
	']':  Token{Type: RBrack},
	'(':  Token{Type: LParen},
	')':  Token{Type: RParen},
	'=':  Token{Type: Equals},
	'&':  Token{Type: Ampersand},
	'!':  Token{Type: Bang},
	'\'': Token{Type: Rune},
	'_':  Token{Type: Underscore},
	':':  Token{Type: Colon},
	';':  Token{Type: Semicolon},
	'<':  Token{Type: LAlligator},
	'>':  Token{Type: RAlligator},
	'.':  Token{Type: Dot},
	',':  Token{Type: Comma},
	'`':  Token{Type: Tack},
}

func Tokenize(content string) []Token {
	ret := []Token{}

	cur := ""

	pos := PositionData{1, 1, 1, 1}
	for i := 0; i < len(content); i++ {
		r := rune(content[i])

		if r == '"' || r == '`' {
			if cur != "" {
				word := Token{cur, Word, pos.scol, pos.sline}
				ret = append(ret, word)
				cur = ""
			}

			raw := tokenizeString(&i, r, content, pos)
			ret = append(ret, Token{raw, String, pos.scol, pos.sline})

			pos.col++
			pos.scol = pos.col
			pos.sline = pos.line

			continue
		}

		if r == '\'' {
			if cur != "" {
				word := Token{cur, Word, pos.scol, pos.sline}
				ret = append(ret, word)
				cur = ""
			}
		}

		if r == '\n' {
			if cur != "" {
				word := Token{cur, Word, pos.scol, pos.sline}
				ret = append(ret, word)
				cur = ""
			}

			pos.col = 1
			pos.line++
			pos.scol = pos.col
			pos.sline = pos.line

			continue
		}

		if r == ' ' || r == '\t' || r == '\r' {
			if cur != "" {
				word := Token{cur, Word, pos.scol, pos.sline}
				ret = append(ret, word)
				cur = ""
			}

			//pos.col++
			pos.scol = pos.col
			pos.sline = pos.line

			continue
		}

		if v, ok := symbols[r]; ok {
			v.Raw = string(r)
			v.Col = pos.col
			v.Line = pos.line

			if cur != "" {
				word := Token{cur, Word, pos.scol, pos.sline}
				ret = append(ret, word)
				cur = ""
			}

			ret = append(ret, v)

			pos.col++
			pos.scol = pos.col
			pos.sline = pos.line

			continue
		}

		cur = cur + string(r)
		pos.col++
	}

	ret = append(ret, Token{"", Eof, 0, 0})
	return ret
}

func tokenizeString(i *int, r rune, content string, pos PositionData) string {
	ret := ""

	*i++

	switch r {
	case '"':
		for {
			if content[*i] == '"' {
				break
			}

			if content[*i] == '\n' {
				pos.line++
				pos.col = 1
			}

			ret += string(content[*i])

			*i++
			pos.col++
		}

	case '`':
		for {
			if content[*i] == '`' {
				break
			}

			if content[*i] == '\n' {
				pos.line++
				pos.col = 1
			}

			ret += string(content[*i])

			*i++
			pos.col++
		}
	}

	return ret
}
