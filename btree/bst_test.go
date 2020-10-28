package btree

import "testing"

func TestAddNode(t *testing.T) {
	bst := &BSTree{}
	bst.Add(NewSimple(10), NewSimple(4), NewSimple(9))
	bst.Add(NewSimple(15), NewSimple(2), NewSimple(12))
	printInorder(t, bst.root)
}

func TestDeleteNode(t *testing.T) {
	bst := &BSTree{}
	nine := NewSimple(9)
	bst.Add(nine, NewSimple(4), NewSimple(8))
	fifteen := NewSimple(15)
	bst.Add(fifteen, NewSimple(2), NewSimple(12))
	seventh := NewSimple(17)
	bst.Add(NewSimple(16), seventh)
	bst.Delete(seventh)
	printInorder(t, bst.root)
}
