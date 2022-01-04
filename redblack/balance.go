package redblack

//func balanceTree(root *Node) (newRoot *Node) {
//	bf := root.getBalanceFactor()
//	if bf > 1 {
//		leftNode := root.left
//		if leftNode != nil && leftNode.right == nil {
//			return rightRotate(root)
//		} else {
//			root.left = leftRotate(leftNode)
//			return rightRotate(root)
//		}
//	} else if bf < -1 {
//		rightNode := root.right
//		if rightNode != nil && rightNode.left == nil {
//			return leftRotate(root)
//		} else {
//			root.right = rightRotate(rightNode)
//			return leftRotate(root)
//		}
//	}
//	return root
//}

func leftRotate(root *Node) (newRoot *Node) {
	newRoot = root.right
	newRoot.parent = root.parent

	root.right = newRoot.left
	if root.right != nil {
		root.right.parent = root
	}

	newRoot.left = root
	root.parent = newRoot

	parent := newRoot.parent
	if parent == nil {
		return newRoot
	}
	if parent.left == root {
		parent.left = newRoot
	} else if parent.right == root {
		parent.right = newRoot
	}

	return newRoot
}

func rightRotate(root *Node) (newRoot *Node) {
	newRoot = root.left
	newRoot.parent = root.parent

	root.left = newRoot.right
	if root.left != nil {
		root.left.parent = root
	}
	newRoot.right = root
	root.parent = newRoot

	parent := newRoot.parent
	if parent == nil {
		return newRoot
	}
	if parent.right == root {
		parent.right = newRoot
	} else if parent.left == root {
		parent.left = newRoot
	}
	return newRoot
}
