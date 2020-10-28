package btree

import (
	"strconv"
)

type nodeGetter interface {
	Val() interface{}
	String() string
	Height() int64
	IsLeaf() bool
	Left() Node
	Right() Node
	Parent() Node
}

type nodeSetter interface {
	AddRight(other Node)
	AddLeft(other Node)
	AddParent(other Node)
	SetHeight(h int64)
}

type nodeOperation interface {
	Eq(other Node) bool
	Greater(other Node) bool
	Less(other Node) bool
	Swap(other Node)
}

type Node interface {
	nodeGetter
	nodeSetter
	nodeOperation
}

func NewSimple(val int64) Node {
	return &SimpleNode{
		val:    val,
		height: 1,
	}
}

type SimpleNode struct {
	val    int64
	height int64
	parent *SimpleNode
	left   *SimpleNode
	right  *SimpleNode
}

func (n *SimpleNode) SetHeight(h int64) {
	n.height = h
}

func (n *SimpleNode) Height() int64 {
	return n.height
}

func (n *SimpleNode) Val() interface{} {
	return n.val
}

func (n *SimpleNode) Eq(other Node) bool {
	val, ok := other.Val().(int64)
	if !ok {
		return false
	}

	return val == n.val && n.height == other.Height()
}

func (n *SimpleNode) Greater(other Node) bool {
	val, ok := other.Val().(int64)
	if !ok {
		return false
	}

	return n.val > val
}

func (n *SimpleNode) Less(other Node) bool {
	val, ok := other.Val().(int64)
	if !ok {
		return false
	}

	return n.val < val
}

func (n *SimpleNode) Left() Node {
	if n.left == nil {
		return nil
	}
	return n.left
}

func (n *SimpleNode) Right() Node {
	if n.right == nil {
		return nil
	}
	return n.right
}

func (n *SimpleNode) Parent() Node {
	if n.parent == nil {
		return nil
	}
	return n.parent
}

func (n *SimpleNode) AddRight(other Node) {
	if other == nil {
		n.right = nil
		return
	}

	sNode, ok := other.(*SimpleNode)
	if !ok {
		return
	}
	n.right = sNode
	other.AddParent(n)
}

func (n *SimpleNode) AddLeft(other Node) {
	if other == nil {
		n.left = nil
		return
	}

	sNode, ok := other.(*SimpleNode)
	if !ok {
		return
	}
	n.left = sNode
	other.AddParent(n)
}

func (n *SimpleNode) AddParent(other Node) {
	if other == nil {
		n.parent = nil
		return
	}

	sNode, ok := other.(*SimpleNode)
	if !ok {
		return
	}
	n.parent = sNode
}

func (n *SimpleNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *SimpleNode) Swap(other Node) {
	sNode, ok := other.(*SimpleNode)
	if !ok {
		return
	}
	nVal := n.val
	n.val = sNode.val
	sNode.val = nVal
}

func (n *SimpleNode) String() string {
	return strconv.Itoa(int(n.val)) + ":" + strconv.Itoa(int(n.height))
}
