package main

import "fmt"

const (
	// B 树的最小度数
	minDegree = 2
)

// B 树的节点
type node struct {
	isLeaf bool    // 是否为叶子节点
	keys   []int   // 节点的键值
	childs []*node // 节点的子节点
	n      int     // 节点包含的键值数量
}

// 创建一个 B 树节点
func newNode(leaf bool) *node {
	return &node{
		isLeaf: leaf,
		keys:   make([]int, 2*minDegree-1),
		childs: make([]*node, 2*minDegree),
		n:      0,
	}
}

// 在 B 树中查找一个键值
func (n *node) search(k int) (bool, int) {
	i := 0
	for i < n.n && k > n.keys[i] {
		i++
	}
	if i < n.n && k == n.keys[i] {
		return true, i
	}
	if n.isLeaf {
		return false, -1
	}
	return n.childs[i].search(k)
}

// 将一个键值插入到 B 树中
func (n *node) insert(k int) *node {
	if n.n == 2*minDegree-1 { // 节点已满，需要拆分
		newNode := newNode(false)
		newNode.childs[0] = n
		newNode.splitChild(0, n)
		newNode.insertNonFull(k)
		return newNode
	} else {
		n.insertNonFull(k)
		return n
	}
}

// 将一个非满节点插入一个键值
func (n *node) insertNonFull(k int) {
	i := n.n - 1
	if n.isLeaf {
		for i >= 0 && n.keys[i] > k {
			n.keys[i+1] = n.keys[i]
			i--
		}
		n.keys[i+1] = k
		n.n++
	} else {
		for i >= 0 && n.keys[i] > k {
			i--
		}
		if n.childs[i+1].n == 2*minDegree-1 { // 子节点已满，需要拆分
			n.splitChild(i+1, n.childs[i+1])
			if n.keys[i+1] < k {
				i++
			}
		}
		n.childs[i+1].insertNonFull(k)
	}
}

// 拆分子节点
func (n *node) splitChild(i int, y *node) {
	z := newNode(y.isLeaf)
	z.n = minDegree - 1

	// 将 y 的后半部分移到 z 中
	for j := 0; j < minDegree-1; j++ {
		z.keys[j] = y.keys[j+minDegree]
	}
	if !y.isLeaf {
		for j := 0; j < minDegree; j++ {
			z.childs[j] = y.childs[j+minDegree]
		}
	}
	// 更新 y 的键值数量
	y.n = minDegree - 1

	// 将 z 插入为 n 的子节点
	for j := n.n; j >= i+1; j-- {
		n.childs[j+1] = n.childs[j]
	}
	n.childs[i+1] = z

	// 将 y 的中间键值插入到 n 中
	for j := n.n - 1; j >= i; j-- {
		n.keys[j+1] = n.keys[j]
	}
	n.keys[i] = y.keys[minDegree-1]
	n.n++
}

// 删除 B 树中的一个键值
func (n *node) delete(k int) *node {
	i := 0
	for i < n.n && n.keys[i] < k {
		i++
	}

	if i < n.n && n.keys[i] == k { // 找到要删除的键值
		if n.isLeaf { // 如果在叶子节点，直接删除
			n.deleteFromLeaf(i)
			return n
		} else {                            // 如果不在叶子节点，则需要将其子节点的一个键值替换为该键值，并在子节点中删除该键值
			if n.childs[i].n >= minDegree { // 左子节点的键值数量大于等于最小度数
				prevKey := n.getPrevKey(i)
				n.keys[i] = prevKey
				n.childs[i].delete(prevKey)
			} else if n.childs[i+1].n >= minDegree { // 右子节点的键值数量大于等于最小度数
				nextKey := n.getNextKey(i)
				n.keys[i] = nextKey
				n.childs[i+1].delete(nextKey)
			} else { // 如果两个子节点都只有最小数量的键值，则需要合并它们
				n.merge(i)
				n.childs[i].delete(k)
			}
		}
	} else { // 如果在非叶子节点，则需要递归删除
		if n.isLeaf {
			return n
		}
		if n.childs[i].n < minDegree {
			n.fill(i)
		}
		n.childs[i].delete(k)
	}
	return n
}

// 从叶子节点中删除一个键值
func (n *node) deleteFromLeaf(i int) {
	for j := i + 1; j < n.n; j++ {
		n.keys[j-1] = n.keys[j]
	}
	n.n--
}

// 获取前一个子节点中的最后一个键值
func (n *node) getPrevKey(i int) int {
	prev := n.childs[i]
	for !prev.isLeaf {
		prev = prev.childs[prev.n]
	}
	return prev.keys[prev.n-1]
}

// 获取后一个子节点中的第一个键值
func (n *node) getNextKey(i int) int {
	next := n.childs[i+1]
	for !next.isLeaf {
		next = next.childs[0]
	}
	return next.keys[0]
}

// 合并 i 和 i+1 节点，其中 i+1 节点的键值数量小于最小度数
func (n *node) merge(i int) {
	child := n.childs[i]
	sibling := n.childs[i+1]
	child.keys[minDegree-1] = n.keys[i]
	// 将 i+1 节点的所有键值和子节点移动到 i 节点
	for j := 0; j < sibling.n; j++ {
		child.keys[j+minDegree] = sibling.keys[j]
	}
	if !sibling.isLeaf {
		for j := 0; j <= sibling.n; j++ {
			child.childs[j+minDegree] = sibling.childs[j]
		}
	}
	child.n += sibling.n + 1

	// 将 n 中 i 之后的所有键值和子节点移动一个位置
	for j := i + 1; j < n.n; j++ {
		n.keys[j-1] = n.keys[j]
	}
	for j := i + 2; j <= n.n; j++ {
		n.childs[j-1] = n.childs[j]
	}
	n.n--

	// 释放 i+1 节点的内存
	sibling = nil
}

// 在 i 节点中填充缺少的键值
func (n *node) fill(i int) {
	if i != 0 && n.childs[i-1].n >= minDegree {
		n.borrowFromPrev(i)
	} else if i != n.n && n.childs[i+1].n >= minDegree {
		n.borrowFromNext(i)
	} else {
		if i != n.n {
			n.merge(i)
		} else {
			n.merge(i - 1)
		}
	}
}

// 从前一个子节点中借一个键值
func (n *node) borrowFromPrev(i int) {
	child := n.childs[i]
	prev := n.childs[i-1]
	// 将 i 节点中的键值后移一位
	for j := child.n - 1; j >= 0; j-- {
		child.keys[j+1] = child.keys[j]
	}
	if !child.isLeaf {
		// 将 i 节点中的子节点后移一位
		for j := child.n; j >= 0; j-- {
			child.childs[j+1] = child.childs[j]
		}
	}
	// 将从前一个子节点中借来的键值插入到 i 节点的开头
	child.keys[0] = n.keys[i-1]
	if !child.isLeaf {
		child.childs[0] = prev.childs[prev.n]
	}
	// 更新 i 节点和前一个子节点的键值数量
	child.n++
	n.keys[i-1] = prev.keys[prev.n-1]
	prev.n--
}

// 从后一个子节点中借一个键值
func (n *node) borrowFromNext(i int) {
	child := n.childs[i]
	next := n.childs[i+1]
	// 将从后一个子节点中借来的键值插入到 i 节点的末尾
	child.keys[child.n] = n.keys[i]
	if !child.isLeaf {
		child.childs[child.n+1] = next.childs[0]
	}
	child.n++
	// 将从后一个子节点中删除的键值和子节点前移
	n.keys[i] = next.keys[0]
	for j := 1; j < next.n; j++ {
		next.keys[j-1] = next.keys[j]
	}
	if !next.isLeaf {
		for j := 1; j <= next.n; j++ {

			next.childs[j-1] = next.childs[j]
		}
	}
	next.n--
}

// 合并 i 和 i+1 子节点
func (n *node) merge(i int) {
	child := n.childs[i]
	sibling := n.childs[i+1]
	// 将 i+1 节点的键值和子节点移动到 i 节点的末尾
	child.keys[minDegree-1] = n.keys[i]
	for j := 0; j < sibling.n; j++ {
		child.keys[j+minDegree] = sibling.keys[j]
	}
	if !child.isLeaf {
		for j := 0; j <= sibling.n; j++ {
			child.childs[j+minDegree] = sibling.childs[j]
		}
	}
	child.n += sibling.n + 1

	// 将 n 中 i 之后的所有键值和子节点前移
	for j := i + 1; j < n.n; j++ {
		n.keys[j-1] = n.keys[j]
	}
	for j := i + 2; j <= n.n; j++ {
		n.childs[j-1] = n.childs[j]
	}
	n.n--

	// 释放 i+1 节点的内存
	sibling = nil

	if n.n == 0 {
		// 如果根节点没有子节点了，则将其指向子节点
		n = child
	}
}

// 删除键值
func (t *Tree) deleteKey(key int) {
	if t.root == nil {
		// 如果树为空，则不做任何操作
		return
	}
	t.root.delete(key)
	if t.root.n == 0 {
		// 如果根节点没有键值了，则将其指向其唯一的子节点
		if !t.root.isLeaf {
			t.root = t.root.childs[0]
		} else {
			t.root = nil
		}
	}
}
