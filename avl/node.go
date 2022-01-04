package avl

type Node struct {
	element     Element
	height      int
	left, right *Node
}

func NewNode(element Element) *Node {
	return &Node{
		element: element,
		height:  1,
	}
}

func (node *Node) getHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) getBalanceFactor() int {
	return node.left.getHeight() - node.right.getHeight()
}

func (node *Node) updateHeight() {
	var max int
	if node.left.getHeight() > node.right.getHeight() {
		max = node.left.getHeight() + 1
	} else {
		max = node.right.getHeight() + 1
	}
	node.height = max
}

type Element interface {
	CompareTo(Element) int
}
