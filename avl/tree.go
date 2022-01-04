package avl

import "fmt"

type Tree interface {
	Insert(key Element)
	Delete(key Element)
	InOrderTraverse()
}

type tree struct {
	root *Node
}

func NewTree() Tree {
	return &tree{}
}

func (t *tree) Insert(e Element) {
	newRoot := insert(t.root, e)
	t.root = newRoot
}

func (t *tree) Delete(key Element) {
	newRoot, isDeleted := delete(t.root, key)
	if !isDeleted {
		return
	}
	t.root = newRoot
}

func (t *tree) InOrderTraverse() {
	inOrderTraverse(t.root)
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

func insert(root *Node, key Element) (newRoot *Node) {
	if root == nil {
		newRoot = NewTreeNode(key)
		return newRoot
	}
	compared := root.element.CompareTo(key)
	if compared < 0 {
		root.right = insert(root.right, key)
	} else if compared > 0 {
		root.left = insert(root.left, key)
	} else {
		return root
	}
	root.updateHeight()
	newRoot = balanceTree(root)
	return newRoot
}

func delete(root *Node, key Element) (newRoot *Node, isDeleted bool) {
	if root == nil {
		return root, false
	}
	compared := root.element.CompareTo(key)
	if compared == 0 {
		if root.left == nil && root.right == nil {
			return nil, true
		} else if root.left == nil {
			return root.right, true
		} else if root.right == nil {
			return root.left, true
		}
		max := greatest(root.left)
		root.element = max.element
		root.left, isDeleted = delete(root.left, key)
	} else if compared < 0 {
		root.right, isDeleted = delete(root.right, key)
	} else {
		root.left, isDeleted = delete(root.left, key)
	}
	root.updateHeight()
	return balanceTree(root), isDeleted
}

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

func greatest(node *Node) (max *Node) {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return greatest(node.right)
}

func inOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.element)
	inOrderTraverse(node.left)
	inOrderTraverse(node.right)
}
