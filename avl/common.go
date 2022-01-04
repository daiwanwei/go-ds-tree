package avl

func greatest(node *Node) (max *Node) {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return greatest(node.right)
}
