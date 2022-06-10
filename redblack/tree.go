package redblack

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
	node := NewNode(e)
	newRoot := insert(t.root, node)
	t.root = newRoot
	t.fixUpForInsert(node)
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

func insert(root *Node, node *Node) (newRoot *Node) {
	if root == nil {
		return node
	}
	compared := root.element.CompareTo(node.element)
	if compared < 0 {
		root.right = insert(root.right, node)
		root.right.parent = root
	} else if compared > 0 {
		root.left = insert(root.left, node)
		root.left.parent = root
	}
	return root
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
	return root, isDeleted
}

func (t *tree) fixUpForInsert(node *Node) {
	tr := node
	for tr != nil && tr.parent != nil && tr.parent.color == Red {
		parent := tr.parent
		grandParent := parent.parent
		if grandParent == nil {
			return
		}
		var uncle *Node
		if grandParent.left == parent {
			uncle = grandParent.right
			if uncle != nil && uncle.color == Red {
				uncle.color = Black
				parent.color = Black
				grandParent.color = Red
				tr = grandParent
			} else {
				if tr == parent.right {
					newParent := leftRotate(parent)
					newParent.color = Black
					grandParent.color = Red
					newGrandParent := rightRotate(grandParent)
					tr = newGrandParent
				} else {
					parent.color = Black
					grandParent.color = Red
					newGrandParent := rightRotate(grandParent)
					tr = newGrandParent
				}
			}
		} else {
			uncle = grandParent.left
			if uncle != nil && uncle.color == Red {
				uncle.color = Black
				parent.color = Black
				grandParent.color = Red
				tr = grandParent
			} else {
				if tr == parent.left {
					newParent := rightRotate(parent)
					newParent.color = Black
					grandParent.color = Red
					newGrandParent := leftRotate(grandParent)
					tr = newGrandParent
				} else {
					parent.color = Black
					grandParent.color = Red
					newGrandParent := leftRotate(grandParent)
					tr = newGrandParent
				}
			}
		}
		if grandParent == t.root {
			t.root = tr
		}
	}
	if t.root != nil && t.root.color == Red {
		t.root.color = Black
	}
}

func (t *tree) fixUpForDelete(node *Node) {
	tr := node
	for tr != nil && tr.parent != nil && tr.parent.color == Red {
		parent := tr.parent
		grandParent := parent.parent
		if grandParent == nil {
			return
		}
		var uncle *Node
		if grandParent.left == parent {
			uncle = grandParent.right
			if uncle != nil && uncle.color == Red {
				uncle.color = Black
				parent.color = Black
				grandParent.color = Red
				tr = grandParent
			} else {
				if tr == parent.right {
					newParent := leftRotate(parent)
					newParent.color = Black
					grandParent.color = Red
					newGrandParent := rightRotate(grandParent)
					tr = newGrandParent
				} else {
					parent.color = Black
					grandParent.color = Red
					newGrandParent := rightRotate(grandParent)
					tr = newGrandParent
				}
			}
		} else {
			uncle = grandParent.left
			if uncle != nil && uncle.color == Red {
				uncle.color = Black
				parent.color = Black
				grandParent.color = Red
				tr = grandParent
			} else {
				if tr == parent.left {
					newParent := rightRotate(parent)
					newParent.color = Black
					grandParent.color = Red
					newGrandParent := leftRotate(grandParent)
					tr = newGrandParent
				} else {
					parent.color = Black
					grandParent.color = Red
					newGrandParent := leftRotate(grandParent)
					tr = newGrandParent
				}
			}
		}
		if grandParent == t.root {
			t.root = tr
		}
	}
	if t.root != nil && t.root.color == Red {
		t.root.color = Black
	}
}
