package bst

type Node struct {
	element     Element
	left, right *Node
}

func NewTreeNode(element Element) *Node {
	return &Node{
		element: element,
	}
}

type Element interface {
	CompareTo(Element) int
}
