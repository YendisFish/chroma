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

type BinOp struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
	Left     Expression
	Right    Expression
	Operator string
}

func (a *BinOp) Children() []Node       { return a.children }
func (a *BinOp) Parent() Node           { return a.parent }
func (a *BinOp) Append(node Node)       { a.children = append(a.children, node) }
func (a *BinOp) CreateParent(node Node) { a.parent = node }
func (a *BinOp) Line() int              { return a.line }
func (a *BinOp) Filename() string       { return a.filename }
func (a *BinOp) Col() int               { return a.col }

type VarRef struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
	Name     string
	IndexRef []Expression
}

func (a *VarRef) Children() []Node       { return a.children }
func (a *VarRef) Parent() Node           { return a.parent }
func (a *VarRef) Append(node Node)       { a.children = append(a.children, node) }
func (a *VarRef) CreateParent(node Node) { a.parent = node }
func (a *VarRef) Line() int              { return a.line }
func (a *VarRef) Filename() string       { return a.filename }
func (a *VarRef) Col() int               { return a.col }

type FuncRef struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
	Name     string
	IndexRef []Expression
}

func (a *FuncRef) Children() []Node       { return a.children }
func (a *FuncRef) Parent() Node           { return a.parent }
func (a *FuncRef) Append(node Node)       { a.children = append(a.children, node) }
func (a *FuncRef) CreateParent(node Node) { a.parent = node }
func (a *FuncRef) Line() int              { return a.line }
func (a *FuncRef) Filename() string       { return a.filename }
func (a *FuncRef) Col() int               { return a.col }
