package linklist

import (
	"math"
	"math/rand"
	"time"
)

type Comparer interface {
	Less(other Comparer) bool
	Greater(other Comparer) bool
	Eq(other Comparer) bool
}

type SkipNode struct {
	Val     Comparer
	forward []*SkipNode
	Level   int32
}

func (node *SkipNode) Forward(level int) *SkipNode {
	if level >= int(node.Level) {
		return nil
	}

	return node.forward[level]
}

func (node *SkipNode) InitForward(level int32) {
	node.forward = make([]*SkipNode, level)
}

type SkipList struct {
	p         float32
	len       int64
	head      *SkipNode
	tail      *SkipNode
	initLevel int32
}

func (sl *SkipList) Insert(node *SkipNode) {
	var (
		resetHead = false
		resetTail = false
	)

	if sl.head == nil {
		node.InitForward(sl.maxLevel())
		sl.head = node
		sl.tail = node
		return
	}

	prevNodesLevel := sl.findPrevNodes(node)

	nodeLevel := sl.randLevel()
	node.InitForward(nodeLevel)

	if prevNodesLevel[0].Forward(0) == nil {
		resetTail = true
	}

	if prevNodesLevel[0] == nil {
		resetHead = true
		nodeLevel = int32(len(sl.head.forward))
	}

	for i := 0; i < int(nodeLevel); i++ {
		if prevNodesLevel[i] == nil {
			node.forward[i] = sl.head
			resetHead = true
			continue
		}
		currentNext := prevNodesLevel[i].forward[i]
		prevNodesLevel[i].forward[i] = node
		node.forward[i] = currentNext
	}

	if resetHead {
		sl.head = node
	}
	if resetTail {
		sl.tail = node
	}
	sl.len++
	sl.checkExpand()
}

func (sl *SkipList) findPrevNodes(node *SkipNode) []*SkipNode {
	var (
		updateNodes = make([]*SkipNode, sl.maxLevel())
		current     = sl.head
	)

	for i := sl.maxLevel() - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].Val.Less(node.Val) {
			current = current.forward[i]
		}
		if current.Val.Less(node.Val) {
			updateNodes[i] = current
		}
	}

	return updateNodes
}

// randLevel 隨機決定層級
//  使用丟銅板方式，正面則++
func (sl *SkipList) randLevel() int32 {
	var level int32 = 1
	rand.Seed(time.Now().Unix())

	for rand.Float32() < sl.p && level < sl.maxLevel() {
		level++
	}
	return level
}

func (sl *SkipList) maxLevel() int32 {
	optimalLevel := int32(math.Sqrt(float64(sl.len)))
	if optimalLevel < sl.initLevel {
		return sl.initLevel
	}

	return optimalLevel
}

func (sl *SkipList) checkExpand() {
	currMax := int32(len(sl.head.forward))
	if currMax < sl.maxLevel() {
		diff := sl.maxLevel() - currMax
		for i := 1; i <= int(diff); i++ {
			sl.head.forward = append(sl.head.forward, sl.tail)
		}
	}
}
