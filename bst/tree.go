package bst

import "fmt"

type Tree interface {
	Insert(key Element)
	Delete(key Element)
	InOrderTraverse()
	PreOrderTraverse()
	PostOrderTraverse()
}

type tree struct {
	root *Node
}

func NewTree() Tree {
	return &tree{}
}

func (t *tree) Insert(e Element) {
	newNode := NewNode(e)
	if t.root == nil {
		t.root = newNode
		return
	}
	insert(t.root, newNode)
}

func (t *tree) InOrderTraverse() {
	inOrderTraverse(t.root)
}

func (t *tree) PreOrderTraverse() {
	preOrderTraverse(t.root)
}

func (t *tree) PostOrderTraverse() {
	postOrderTraverse(t.root)
}

func (t *tree) MinNode() (node *Node) {
	if t.root == nil {
		return
	}
	node = t.root
	for {
		if node.left == nil {
			return node
		}
		node = node.left
	}
}

func (t *tree) MaxNode() (node *Node) {
	if t.root == nil {
		return
	}
	node = t.root
	for {
		if node.right == nil {
			return node
		}
		node = node.right
	}
}

func (t *tree) SearchNode(e Element) (node *Node) {
	return searchNode(t.root, e)
}

func (t *tree) Delete(e Element) {
	delete(t.root, e)
}

func (t *tree) String() {
	fmt.Println("------------------------------------------------")
	string(t.root, 0)
	fmt.Println("------------------------------------------------")
}

func string(node *Node, level int) {
	if node == nil {
		return
	}
	format := ""
	for i := 0; i < level; i++ {
		format += " "
	}
	format += "---[ "
	string(node.right, level+1)
	fmt.Printf(format+"%d\n", node.element)
	string(node.left, level+1)
}

func delete(node *Node, e Element) (newNode *Node) {
	if node == nil {
		return
	}
	compared := node.element.CompareTo(e)
	if compared > 0 {
		newLeft := delete(node.left, e)
		node.left = newLeft
		return node
	} else if compared < 0 {
		newRight := delete(node.right, e)
		node.right = newRight
		return node
	}

	if node.right == nil && node.left == nil {
		return nil
	} else if node.right == nil {
		return node.left
	} else if node.left == nil {
		return node.right
	}
	var maxNodeInLeft, traverseNode *Node
	traverseNode = node.left
	for {
		if traverseNode != nil && traverseNode.right != nil {
			maxNodeInLeft = traverseNode
			traverseNode = traverseNode.right
		} else {
			break
		}
	}
	newNode = NewNode(maxNodeInLeft.element)
	newNode.right = node.right
	newNode.left = delete(node.left, maxNodeInLeft.element)
	return

}

func insert(node, newNode *Node) {
	compared := node.element.CompareTo(newNode.element)
	if compared > 0 {
		if node.left == nil {
			node.left = newNode
			return
		} else {
			insert(node.left, newNode)
		}
	} else if compared < 0 {
		if node.right == nil {
			node.right = newNode
			return
		} else {
			insert(node.right, newNode)
		}
	}
}
