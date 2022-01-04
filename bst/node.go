package bst

type Node struct {
	element     Element
	left, right *Node
}

func NewNode(element Element) *Node {
	return &Node{
		element: element,
	}
}

type Element interface {
	CompareTo(Element) int
}
