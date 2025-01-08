package lexer

type Token struct {
	Raw  string
	Type TokenType
	Col  int
	Line int
}

type TokenType int

const (
	Add TokenType = iota
	Subtract
	FSlash
	Star
	BSlash
	LBrace
	RBrace
	LBrack
	RBrack
	LParen
	RParen
	Equals
	Ampersand
	Bang
	String
	Rune
	Word
	Underscore
	Colon
	Semicolon
	LAlligator
	RAlligator
	Dot
	Comma
	Tack
	Eof

	//there will be more for binary operators
)

/*func NewToken(raw string, col int, line int, tp TokenType) Token {
	return Token{raw, tp, col, line}
}*/

type PositionData struct {
	scol  int
	sline int
	col   int
	line  int
}
