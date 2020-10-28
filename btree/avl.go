package btree

type AVLTree struct {
	*BSTree
}

func (t *AVLTree) Add(nodes ...Node) {
	for i := range nodes {
		if t.root == nil {
			t.root = nodes[i]
			continue
		}
		t.add(t.root, nodes[i])
		t.rotate(nodes[i])
	}
}

func (t *AVLTree) Delete(node Node) {
	t.delete(node)
	t.rotate(node)
}

func (t *AVLTree) rotate(node Node) {
	if node == nil || node.Parent() == nil || node.Parent().Parent() == nil {
		return
	}

	target := node.Parent().Parent()
	lH, rH := getLRHeghts(target)
	diff := abs(rH - lH)
	rotated := false
	if diff >= 2 && lH > rH {
		t.rotateRight(target)
		rotated = true
	}

	if diff >= 2 && lH < rH {
		t.rotateLeft(target)
		rotated = true
	}

	if rotated && target.Parent().Parent() != nil {
		updateParentHeight(target.Parent())
	}
}

func (t *AVLTree) rotateLeft(node Node) {
	var (
		newParent, newRoot, newLeftRight Node
	)

	if node == nil {
		return
	}

	newParent = node.Parent()
	newRoot = node.Right()
	if newRoot != nil {
		newLeftRight = node.Right().Left()
		newRoot.AddParent(newParent)
		if newParent != nil {
			if newParent.Less(newRoot) {
				newParent.AddRight(newRoot)
			} else {
				newParent.AddLeft(newRoot)
			}
		}
		newRoot.AddLeft(node)
		node.AddParent(newRoot)
		node.AddRight(newLeftRight)
		updateHeight(node)
		updateHeight(newRoot)
	}

	if newLeftRight != nil {
		newLeftRight.AddParent(node)
	}

	if newRoot.Parent() == nil {
		t.root = newRoot
	}
}

func (t *AVLTree) rotateRight(node Node) {
	var (
		newParent, newRoot, newRightLeft Node
	)

	if node == nil {
		return
	}

	newParent = node.Parent()
	newRoot = node.Left()
	if newRoot != nil {
		newRightLeft = node.Left().Right()
		newRoot.AddParent(newParent)
		if newParent != nil {
			if newParent.Less(newRoot) {
				newParent.AddRight(newRoot)
			} else {
				newParent.AddLeft(newRoot)
			}
		}
		newRoot.AddRight(node)
		node.AddParent(newRoot)
		node.AddLeft(newRightLeft)
		updateHeight(node)
		updateHeight(newRoot)
	}

	if newRightLeft != nil {
		newRightLeft.AddParent(node)
	}

	if newRoot.Parent() == nil {
		t.root = newRoot
	}
}
