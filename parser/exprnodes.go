package parser

type Expression interface {
	Node
}

type Literal struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
	Raw      string
	Type     LiteralType
}

func (a *Literal) Children() []Node       { return a.children }
func (a *Literal) Parent() Node           { return a.parent }
func (a *Literal) Append(node Node)       { a.children = append(a.children, node) }
func (a *Literal) CreateParent(node Node) { a.parent = node }
func (a *Literal) Line() int              { return a.line }
func (a *Literal) Filename() string       { return a.filename }
func (a *Literal) Col() int               { return a.col }

type LiteralType int

const (
	LByte LiteralType = iota
	LNum
	LString
	LRune
	LTrue
	LFalse
)
