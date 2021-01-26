package tree

import (
	"testing"
)

func generateTree() *BinarySearchTree {
	tree := &BinarySearchTree{}
	tree.Insert(4)
	tree.Insert(1)
	tree.Insert(22)
	tree.Insert(699)
	tree.Insert(88)
	return tree
}

func TestFind(t *testing.T) {
	tree := generateTree()
	got := tree.Find(699)

	if got == nil {
		t.Errorf("TestFind failed, expected: %d", 699)
	}
}
