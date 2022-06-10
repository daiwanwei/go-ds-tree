package bst

import (
	"fmt"
	"testing"
)

type element int

func (e element) CompareTo(other Element) int {
	if e == other.(element) {
		return 0
	} else if e > other.(element) {
		return 1
	} else {
		return -1
	}
}

func TestTree(t *testing.T) {
	tree := NewTree()
	tree.Insert(element(1))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(3))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(2))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(5))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(1))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Delete(element(5))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
}
