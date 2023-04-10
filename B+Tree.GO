package main

import "fmt"

const minDegree = 2

type node struct {
	isLeaf   bool      // 是否为叶子节点
	n        int       // 当前节点包含的键值个数
	keys     []int     // 键值
	pointers []pointer // 指向子节点或数据的指针
	parent   *node     // 父节点
	next     *node     // 下一个叶子节点
}

type pointer struct {
	isLeaf bool // 是否为叶子节点
	node   *node
	value  int
}

type Tree struct {
	root *node // 根节点
}

// 创建一个新节点
func newNode(isLeaf bool) *node {
	return &node{
		isLeaf:   isLeaf,
		n:        0,
		keys:     make([]int, 2*minDegree-1),
		pointers: make([]pointer, 2*minDegree),
		parent:   nil,
		next:     nil,
	}
}

// 在节点中查找键值应该插入的位置
func (n *node) findIndex(key int) int {
	l, r := 0, n.n-1
	for l <= r {
		m := (l + r) / 2
		if n.keys[m] == key {
			return m
		}
		if n.keys[m] < key {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

// 在节点中查找指定键值的位置，如果找不到则返回 nil
func (n *node) find(key int) *pointer {
	i := n.findIndex(key)
	if i < n.n && n.keys[i] == key {
		return &n.pointers[i]
	}
	if n.isLeaf {
		return nil
	}
	return n.pointers[i].node.find(key)
}

// 将新的子节点插入到节点中的指定位置
func (n *node) insertChildAt(child *node, i int) {
	for j := n.n; j > i; j-- {
		n.pointers[j] = n.pointers[j-1]
	}
	n.pointers[i].node = child
	child.parent = n
	n.n++
}

// 分裂节点
func (n *node) split() (*node, int) {
	right := newNode(n.isLeaf)

	// 计算分裂点位置
	mid := n.n / 2

	// 将 n 中的右半部分键值和子节点移动到 right 中
	right.n = n.n - mid - 1
	for i := 0; i < right.n; i++ {
		right.keys[i] = n.keys[mid+i+1]
		right.pointers[i] = n.pointers[mid+i+1]
	}
	right.pointers[right.n] = n.pointers[n.n]

	// 更新 n 的键值个数
	n.n = mid

	return right, n.keys[mid]
}

// 插入键值
func (n *node) insert(key int, value int) {
	// 找到应该插入的位置
	i := n.findIndex(key)

	// 如果键值已存在，则更新其值
	if i < n.n && n.keys[i] == key {
		n.pointers[i].value = value
		return
	}

	// 如果是叶子节点，则直接将新键值插入到对应位置
	if n.isLeaf {
		for j := n.n; j > i; j-- {
			n.keys[j] = n.keys[j-1]
			n.pointers[j] = n.pointers[j-1]
		}
		n.keys[i] = key
		n.pointers[i] = pointer{isLeaf: true, value: value}
		n.n++

		// 如果插入后当前节点的键值个数超过了阈值，则需要进行分裂操作
		if n.n > 2*minDegree-1 {
			right, midKey := n.split()
			if n.parent == nil {
				n.parent = newNode(false)
				n.parent.insertChildAt(n, 0)
			}
			i := n.parent.findIndex(midKey)
			n.parent.insertChildAt(right, i+1)
			if right.pointers[0].isLeaf {
				right.next = n.next
				n.next = right
			}
		}
	} else {
		// 如果不是叶子节点，则在对应子节点中递归插入键值
		n.pointers[i].node.insert(key, value)

		// 如果递归调用后子节点分裂了，则需要将其分裂后的右半部分插入到当前节点中
		if n.pointers[i].node.n > 2*minDegree-1 {
			right, midKey := n.pointers[i].node.split()
			n.insertChildAt(right, i+1)
			if n.n > 2*minDegree-1 {
				if n.parent == nil {
					n.parent = newNode(false)
					n.parent.insertChildAt(n, 0)
				}
				i := n.parent.findIndex(midKey)
				n.parent.insertChildAt(right, i+1)
			}
		}
	}
}

// 从节点中删除键值
func (n *node) delete(key int) {
	// 找到应该删除的位置
	i := n.findIndex(key)
	// 如果键值在节点中，则进行删除操作
	if i < n.n && n.keys[i] == key {
		if n.isLeaf {
			// 如果是叶子节点，则直接删除该键值
			for j := i; j < n.n-1; j++ {
				n.keys[j] = n.keys[j+1]
				n.pointers[j] = n.pointers[j+1]
			}
			n.n--
		} else {
			// 如果不是叶子节点，则需要将该键值替换为其后继节点的键值，再在其后继节点中递归删除
			nextNode := n.pointers[i+1].node
			for !nextNode.isLeaf {
				nextNode = nextNode.pointers[0].node
			}
			n.keys[i] = nextNode.keys[0]
			nextNode.delete(nextNode.keys[0])
		}
	} else {
		// 如果键值不在节点中，则在对应子节点中递归删除
		if n.isLeaf {
			return
		}
		n.pointers[i].node.delete(key)
	}

	// 如果删除后节点的键值个数小于阈值
	if n.n < minDegree && n.parent != nil {
		// 先找到当前节点在父节点中的位置
		i := n.parent.findIndex(n.keys[0])

		// 如果当前节点的前一个兄弟节点存在且键值个数大于阈值，则从其处借一个键值
		if i > 0 && n.parent.pointers[i-1].node.n > minDegree-1 {
			leftSibling := n.parent.pointers[i-1].node
			lastKeyIndex := leftSibling.n - 1

			// 将父节点中对应位置的键值下移至当前节点中
			n.keys = append([]int{0}, n.keys...)
			copy(n.keys[1:], n.keys[:n.n])
			n.keys[0] = n.parent.keys[i-1]
			n.n++

			// 将左兄弟节点中最后一个键值上移至父节点中
			n.pointers = append([]pointer{{}}, n.pointers...)
			copy(n.pointers[1:], n.pointers[:n.n])
			n.pointers[0] = leftSibling.pointers[lastKeyIndex]
			n.pointers[0].node = leftSibling.pointers[lastKeyIndex].node
			leftSibling.n--
			n.parent.keys[i-1] = leftSibling.keys[lastKeyIndex]
			return
		}

		// 如果当前节点的后一个兄弟节点存在且键值个数大于阈值，则从其处借一个键值
		if i < n.parent.n && n.parent.pointers[i+1].node.n > minDegree-1 {
			rightSibling := n.parent.pointers[i+1].node

			// 将父节点中对应位置的键值下移至当前节点中
			n.keys = append(n.keys, n.parent.keys[i])
			n.n++

			// 将右兄弟节点中第一个键值上移至父节点中
			n.pointers = append(n.pointers, rightSibling.pointers[0])
			copy(n.pointers[n.n:], n.pointers[n-1:])
			n.pointers[n].node = rightSibling.pointers[0].node
			rightSibling.n--
			n.parent.keys[i] = rightSibling.keys[0]

			// 如果右兄弟节点是内部节点，则需要将其第一个子节点的父节点指向当前节点
			if !rightSibling.isLeaf {
				rightSibling.pointers[0] = rightSibling.pointers[1]
				rightSibling.pointers[0].node.parent = n
				rightSibling.pointers = rightSibling.pointers[1:]
			}
			return
		}

		// 如果当前节点的前后兄弟节点都不存在或者键值个数都小于等于阈值，则需要合并节点
		if i > 0 {
			// 先合并前一个兄弟节点
			leftSibling := n.parent.pointers[i-1].node
			lastKeyIndex := leftSibling.n
			leftSibling.keys = append(leftSibling.keys, n.parent.keys[i-1])
			leftSibling.n++
			leftSibling.keys = append(leftSibling.keys, n.keys...)
			leftSibling.pointers = append(leftSibling.pointers, n.pointers...)
			for j := range n.pointers {
				n.pointers[j].node.parent = leftSibling
			}
			n.parent.keys = append(n.parent.keys[:i-1], n.parent.keys[i:]...)
			n.parent.pointers = append(n.parent.pointers[:i], n.parent.pointers[i+1:]...)
			n.parent.n--
			n = leftSibling
		} else {
			// 否则合并后一个兄弟节点
			rightSibling := n.parent.pointers[i+1].node
			n.keys = append(n.keys, n.parent.keys[i])
			n.n++
			n.keys = append(n.keys, rightSibling.keys...)
			n.pointers = append(n.pointers, rightSibling.pointers...)
			for j := range rightSibling.pointers {
				rightSibling.pointers[j].node.parent = n
			}
			n.parent.keys = append(n.parent.keys[:i], n.parent.keys[i+1:]...)
			n.parent.pointers = append(n.parent.pointers[:i+1], n.parent.pointers[i+2:]...)
			n.parent.n--
		}
		// 递归地向上调整
		n.adjust()
	}

}

// Remove 删除键值为key的记录
func (t *BPlusTree) Remove(key int) bool {
	// 查找对应的叶子节点
	node, index := t.searchLeaf(key)
	if node == nil || index == -1 {
		return false
	}
	// 删除节点中的对应键值
	copy(node.keys[index:], node.keys[index+1:])
	copy(node.pointers[index:], node.pointers[index+1:])
	node.n--

	// 调整树结构
	node.adjust()

	return true
}
