package avl

import "fmt"

func inOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.element)
	inOrderTraverse(node.left)
	inOrderTraverse(node.right)
}
