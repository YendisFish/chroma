package parser

//function

type Function struct {
	parent    Node
	children  []Node
	Name      string
	Type      TypeInfo
	Arguments []Variable
}

func (a *Function) Children() []Node       { return a.children }
func (a *Function) Parent() Node           { return a.parent }
func (a *Function) Append(node Node)       { a.children = append(a.children, node) }
func (a *Function) CreateParent(node Node) { a.parent = node }

//

//variable

type Variable struct {
	parent   Node
	children []Node
	Name     string
	Type     TypeInfo
}

func (a *Variable) Children() []Node       { return a.children }
func (a *Variable) Parent() Node           { return a.parent }
func (a *Variable) Append(node Node)       { a.children = append(a.children, node) }
func (a *Variable) CreateParent(node Node) { a.parent = node }

//
