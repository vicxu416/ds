package btree

import "testing"

func printInorder(t *testing.T, node Node) {
	nodes := []Node{}
	Traversal(node, &nodes)
	t.Logf("%s", nodes)

}

func TestAVLAdd(t *testing.T) {
	avl := &AVLTree{BSTree: &BSTree{}}
	avl.Add(NewSimple(10), NewSimple(11), NewSimple(12), NewSimple(13))
	avl.Add(NewSimple(9), NewSimple(8))
	avl.Add(NewSimple(14), NewSimple(15))
	printInorder(t, avl.root)
	t.Logf("root:%s", avl.root)
}

func TestAVLDelete(t *testing.T) {
	avl := &AVLTree{BSTree: &BSTree{}}
	ten := NewSimple(10)
	avl.Add(ten, NewSimple(11), NewSimple(12), NewSimple(13))
	avl.Add(NewSimple(9), NewSimple(8), NewSimple(7))
	avl.Add(NewSimple(14), NewSimple(15))
	avl.Delete(ten)
	printInorder(t, avl.root)
}

func TestRotateRight(t *testing.T) {
	avl := &AVLTree{}

	// node := NewSimple(10)
	// node.AddRight(NewSimple(11))
	// node.AddLeft(NewSimple(9))
	// node.Left().AddLeft(NewSimple(8))
	// avl.rotateRight(node)

	// printInorder(t, node)

	// node = NewSimple(10)
	// node.AddLeft(NewSimple(8))
	// node.Left().AddRight(NewSimple(9))
	// avl.rotateRight(node)
	// printInorder(t, node)

	node := NewSimple(11)
	node2 := NewSimple(10)
	node.AddLeft(node2)
	node.Left().AddLeft(NewSimple(9))
	node.Left().Left().AddLeft(NewSimple(8))
	avl.rotateRight(node2)
	t.Logf("%s", node.Left())
	printInorder(t, node)
}

func TestRotateLeft(t *testing.T) {
	avl := &AVLTree{}
	node := NewSimple(10)
	node.AddRight(NewSimple(11))
	node.Right().AddRight(NewSimple(12))
	node.AddLeft(NewSimple(9))
	avl.rotateLeft(node)
	printInorder(t, node)
}
