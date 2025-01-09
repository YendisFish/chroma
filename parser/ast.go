package parser

type Ast struct {
	Root Node
}

type Node interface {
	Parent() Node
	CreateParent(Node)
	Children() []Node
	Append(node Node)
	//Line() int
	//Filename() string
}

type RootNode struct {
	parent   Node
	children []Node
}

func (a *RootNode) Children() []Node       { return a.children }
func (a *RootNode) Parent() Node           { return a.parent }
func (a *RootNode) Append(node Node)       { a.children = append(a.children, node) }
func (a *RootNode) CreateParent(node Node) { a.parent = node }

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
