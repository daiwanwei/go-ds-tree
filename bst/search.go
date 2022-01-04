package bst

func searchNode(node *Node, e Element) (result *Node) {
	if node == nil {
		return
	}
	compared := node.element.CompareTo(e)
	if compared == 0 {
		return node
	} else if compared > 0 {
		return searchNode(node.left, e)
	} else {
		return searchNode(node.right, e)
	}
}
