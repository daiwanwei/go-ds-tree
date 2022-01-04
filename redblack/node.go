package redblack

type Node struct {
	element             Element
	color               Color
	left, right, parent *Node
}

func NewNode(element Element) *Node {
	return &Node{
		element: element,
		color:   Red,
	}
}

type Element interface {
	CompareTo(Element) int
}

type Color int

const (
	Red Color = iota
	Black
)
