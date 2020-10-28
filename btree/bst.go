package btree

type BSTree struct {
	root Node
}

func (t *BSTree) MaxHeight() int64 {
	return 0
}

func (t *BSTree) traversal(parent Node, nodes *[]Node) {
	if parent == nil {
		return
	}

	if parent.IsLeaf() {
		*nodes = append(*nodes, parent)
		return
	}

	t.traversal(parent.Left(), nodes)
	*nodes = append(*nodes, parent)
	t.traversal(parent.Right(), nodes)
}

func (t *BSTree) Add(nodes ...Node) {
	for i := range nodes {
		if t.root == nil {
			t.root = nodes[i]
			continue
		}

		t.add(t.root, nodes[i])
	}
}

func (t *BSTree) Delete(node Node) {
	if t.root == nil {
		return
	}
	t.delete(node)
	updateParentHeight(node)
}

func (t *BSTree) add(parent, node Node) {
	if parent.Greater(node) || parent.Eq(node) {
		if parent.Left() == nil {
			parent.AddLeft(node)
			updateHeight(parent)
			return
		}
		t.add(parent.Left(), node)
		updateHeight(parent)
		return
	}

	if parent.Right() == nil {
		parent.AddRight(node)
		updateHeight(parent)
		return
	}
	t.add(parent.Right(), node)
	updateHeight(parent)
}

func (t *BSTree) delete(node Node) {
	var parentLeft, parentRight Node

	if node.Parent() != nil {
		parentLeft = node.Parent().Left()
		parentRight = node.Parent().Right()
	}
	if node.IsLeaf() {
		if parentLeft != nil && parentLeft.Eq(node) {
			node.Parent().AddLeft(nil)
		}
		if parentRight != nil && parentRight.Eq(node) {
			node.Parent().AddRight(nil)
		}
		updateHeight(node.Parent())
		node = nil
		return
	}

	var target Node
	if node.Right() == nil {
		target = node.Left()
	}
	if node.Left() == nil {
		target = node.Right()
	}
	if target != nil {
		if parentLeft != nil && parentLeft.Eq(node) {
			node.Parent().AddLeft(target)
		}
		if parentRight != nil && parentRight.Eq(node) {
			node.Parent().AddRight(target)
		}
		target.AddParent(node.Parent())
		node = nil
		return
	}

	rightMin := findRightMin(node)
	node.Swap(rightMin)
	t.delete(rightMin)
	updateHeight(node)
}
