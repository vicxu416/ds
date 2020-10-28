package btree

func Traversal(parent Node, nodes *[]Node) {
	if parent == nil {
		return
	}

	if parent.IsLeaf() {
		*nodes = append(*nodes, parent)
		return
	}

	Traversal(parent.Left(), nodes)
	*nodes = append(*nodes, parent)
	Traversal(parent.Right(), nodes)
}

func findRightMin(node Node) Node {
	target := node.Right()
	if target == nil {
		return nil
	}
	if target.IsLeaf() {
		return target
	}

	for target.Left() != nil {
		target = target.Left()
	}
	return target
}

func findLeftMax(node Node) Node {
	target := node.Left()
	if target == nil {
		return nil
	}
	if target.IsLeaf() {
		return target
	}

	for target.Right() != nil {
		target = target.Right()
	}
	return target
}

func getLRHeghts(node Node) (int64, int64) {
	var lH, rH int64
	if node.Left() != nil {
		lH = node.Left().Height()
	}

	if node.Right() != nil {
		rH = node.Right().Height()
	}
	return lH, rH
}

func updateParentHeight(node Node) {
	parent := node.Parent()

	if parent == nil {
		return
	}
	updateHeight(parent)
	updateParentHeight(parent)
}

func updateHeight(node Node) {
	lH, rH := getLRHeghts(node)
	node.SetHeight(max(lH, rH) + 1)
}

func max(nums ...int64) int64 {
	max := int64(0)

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
