package avl

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

func insert(root *Node, key Element) (newRoot *Node) {
	if root == nil {
		newRoot = NewNode(key)
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
