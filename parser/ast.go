package parser

type Ast struct {
	Root Node
}

type Node interface {
	Parent() Node
	CreateParent(Node)
	Children() []Node
	Append(node Node)
	Line() int
	Filename() string
	Col() int
}

type RootNode struct {
	parent   Node
	children []Node
	line     int
	col      int
	filename string
}

func (a *RootNode) Children() []Node       { return a.children }
func (a *RootNode) Parent() Node           { return a.parent }
func (a *RootNode) Append(node Node)       { a.children = append(a.children, node) }
func (a *RootNode) CreateParent(node Node) { a.parent = node }
func (a *RootNode) Line() int              { return a.line }
func (a *RootNode) Filename() string       { return a.filename }
func (a *RootNode) Col() int               { return a.col }

type TypeInfo struct {
	Name     *string
	Generics []TypeInfo
	Pointers []Ptr
}

func EmptyType() TypeInfo {
	return TypeInfo{nil, nil, nil}
}

type Ptr struct {
	Type  PointerType
	Index *int64
}

type PointerType int

const (
	Array PointerType = iota
	Pointer
	Slice
)
