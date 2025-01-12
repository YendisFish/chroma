package parser

//function

type Function struct {
	parent    Node
	children  []Node
	line      int
	col       int
	filename  string
	Name      string
	Type      TypeInfo
	Arguments []Variable
}

func (a *Function) Children() []Node       { return a.children }
func (a *Function) Parent() Node           { return a.parent }
func (a *Function) Append(node Node)       { a.children = append(a.children, node) }
func (a *Function) CreateParent(node Node) { a.parent = node }
func (a *Function) Line() int              { return a.line }
func (a *Function) Filename() string       { return a.filename }
func (a *Function) Col() int               { return a.col }

//

//variable

type Variable struct {
	parent     Node
	children   []Node
	Name       string
	Type       TypeInfo
	line       int
	col        int
	filename   string
	assignment Expression
}

func (a *Variable) Children() []Node       { return a.children }
func (a *Variable) Parent() Node           { return a.parent }
func (a *Variable) Append(node Node)       { a.children = append(a.children, node) }
func (a *Variable) CreateParent(node Node) { a.parent = node }
func (a *Variable) Line() int              { return a.line }
func (a *Variable) Filename() string       { return a.filename }
func (a *Variable) Col() int               { return a.col }

//
