package bst

import "fmt"

func postOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	if node.left != nil {
		inOrderTraverse(node.left)
	}
	if node.right != nil {
		inOrderTraverse(node.right)
	}
	fmt.Println(node.element)
}

func preOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.element)
	if node.left != nil {
		inOrderTraverse(node.left)
	}
	if node.right != nil {
		inOrderTraverse(node.right)
	}
}

func inOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	if node.left != nil {
		inOrderTraverse(node.left)
	}
	fmt.Println(node.element)
	if node.right != nil {
		inOrderTraverse(node.right)
	}
}
