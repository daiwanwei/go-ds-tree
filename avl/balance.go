package avl

func balanceTree(root *Node) (newRoot *Node) {
	bf := root.getBalanceFactor()
	if bf > 1 {
		leftNode := root.left
		if leftNode != nil && leftNode.right == nil {
			return rightRotate(root)
		} else {
			root.left = leftRotate(leftNode)
			return rightRotate(root)
		}
	} else if bf < -1 {
		rightNode := root.right
		if rightNode != nil && rightNode.left == nil {
			return leftRotate(root)
		} else {
			root.right = rightRotate(rightNode)
			return leftRotate(root)
		}
	}
	return root
}

func leftRotate(root *Node) (newRoot *Node) {
	newRoot = root.right
	root.right = newRoot.left
	newRoot.left = root

	root.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func rightRotate(root *Node) (newRoot *Node) {
	newRoot = root.left
	root.left = newRoot.right
	newRoot.right = root

	root.updateHeight()
	newRoot.updateHeight()
	return newRoot
}
