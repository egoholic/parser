package ast

type Node struct {
	Parsed     []rune
	ParsedWith string
	Children   []*Node
}

func New(vertexName string, parsed []rune) *Node {
	return &Node{
		ParsedWith: vertexName,
		Parsed:     parsed,
		Children:   []*Node{},
	}
}

func (n *Node) AddChild(ch *Node) {
	n.Children = append(n.Children, ch)
}
