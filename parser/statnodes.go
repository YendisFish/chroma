package parser

type Statement interface {
	Node
}

type If struct {
	parent    Node
	children  []Node
	line      int
	col       int
	filename  string
	Condition Expression
	ElseBlock *Else
}

func (a *If) Children() []Node       { return a.children }
func (a *If) Parent() Node           { return a.parent }
func (a *If) Append(node Node)       { a.children = append(a.children, node) }
func (a *If) CreateParent(node Node) { a.parent = node }
func (a *If) Line() int              { return a.line }
func (a *If) Filename() string       { return a.filename }
func (a *If) Col() int               { return a.col }

type Else struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
}

func (a *Else) Children() []Node       { return a.children }
func (a *Else) Parent() Node           { return a.parent }
func (a *Else) Append(node Node)       { a.children = append(a.children, node) }
func (a *Else) CreateParent(node Node) { a.parent = node }
func (a *Else) Line() int              { return a.line }
func (a *Else) Filename() string       { return a.filename }
func (a *Else) Col() int               { return a.col }

type While struct {
	parent    Node
	children  []Node
	line      int
	col       int
	filename  string
	Condition Expression
}

func (a *While) Children() []Node       { return a.children }
func (a *While) Parent() Node           { return a.parent }
func (a *While) Append(node Node)       { a.children = append(a.children, node) }
func (a *While) CreateParent(node Node) { a.parent = node }
func (a *While) Line() int              { return a.line }
func (a *While) Filename() string       { return a.filename }
func (a *While) Col() int               { return a.col }

type Package struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
	Name     string
	Type     PkgType
}

func (a *Package) Children() []Node       { return a.children }
func (a *Package) Parent() Node           { return a.parent }
func (a *Package) Append(node Node)       { a.children = append(a.children, node) }
func (a *Package) CreateParent(node Node) { a.parent = node }
func (a *Package) Line() int              { return a.line }
func (a *Package) Filename() string       { return a.filename }
func (a *Package) Col() int               { return a.col }

type PkgType int

const (
	PkgNormal PkgType = iota
	PkgAllocator
)

type Import struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
	Name     string
}

func (a *Import) Children() []Node       { return a.children }
func (a *Import) Parent() Node           { return a.parent }
func (a *Import) Append(node Node)       { a.children = append(a.children, node) }
func (a *Import) CreateParent(node Node) { a.parent = node }
func (a *Import) Line() int              { return a.line }
func (a *Import) Filename() string       { return a.filename }
func (a *Import) Col() int               { return a.col }

type For struct {
	parent    Node
	children  []Node
	line      int
	col       int
	filename  string
	Var       *Variable
	Cond      Expression
	Operation Expression
}

func (a *For) Children() []Node       { return a.children }
func (a *For) Parent() Node           { return a.parent }
func (a *For) Append(node Node)       { a.children = append(a.children, node) }
func (a *For) CreateParent(node Node) { a.parent = node }
func (a *For) Line() int              { return a.line }
func (a *For) Filename() string       { return a.filename }
func (a *For) Col() int               { return a.col }
