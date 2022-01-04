package redblack

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
	tree.Insert(element(4))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(7))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(12))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(15))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(3))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(5))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(14))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
	tree.Insert(element(18))
	tree.InOrderTraverse()
	fmt.Println("------------------------")
}
