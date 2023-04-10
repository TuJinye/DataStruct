package main

import "fmt"

type node struct {
	val         int
	color       int // 0 为黑色，1 为红色
	left, right *node
	parent      *node
}

type redBlackTree struct {
	root *node
}

// 左旋操作
func (t *redBlackTree) rotateLeft(x *node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// 右旋操作
func (t *redBlackTree) rotateRight(x *node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// 插入操作
func (t *redBlackTree) insert(val int) {
	// 创建新节点
	x := &node{val: val, color: red}

	// 找到新节点的插入位置
	var y *node
	z := t.root
	for z != nil {
		y = z
		if x.val < z.val {
			z = z.left
		} else {
			z = z.right
		}
	}

	// 将新节点插入树中
	x.parent = y
	if y == nil {
		t.root = x
	} else if x.val < y.val {
		y.left = x
	} else {
		y.right = x
	}

	// 红黑树性质的调整
	t.insertFixup(x)
}

// 插入后的红黑树性质的调整
func (t *redBlackTree) insertFixup(x *node) {
	// 父节点为红色，说明违反了红黑树性质，需要进行调整
	for x.parent != nil && x.parent.color == red {
		// 如果新节点的父节点是祖父节点的左孩子
		if x.parent == x.parent.parent.left {
			// 将 y 设为 x 的叔叔节点
			y := x.parent.parent.right

			// 如果 y 是红色的，情况 1
			if y != nil && y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else {
				// 如果 y 是黑色的，情况 2 和情况 3
				if x == x.parent.right {
					x = x.parent
					t.rotateLeft(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				t.rotateRight(x.parent.parent)
			}
		} else {
			// 将 y 设为 x 的叔叔节点
			y := x.parent
			// 如果 y 是红色的，情况 1
			if y != nil && y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else {
				// 如果 y 是黑色的，情况 2 和情况 3
				if x == x.parent.left {
					x = x.parent
					t.rotateRight(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				t.rotateLeft(x.parent.parent)
			}
		}
	}
	// 根节点必须是黑色
	t.root.color = black
}

// 查找最小值节点
func (t *redBlackTree) minNode(x *node) *node {
	for x.left != nil {
		x = x.left
	}
	return x
}

// 删除节点
func (t *redBlackTree) delete(val int) {
	// 找到要删除的节点
	z := t.root
	for z != nil {
		if z.val == val {
			break
		} else if z.val < val {
			z = z.right
		} else {
			z = z.left
		}
	}
	if z == nil {
		return
	}

	var x, y *node

	// 找到替代节点 y
	if z.left == nil || z.right == nil {
		y = z
	} else {
		y = t.minNode(z.right)
	}

	// 找到 y 的子节点 x
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	// 将 x 的父节点指向 y 的父节点
	if x != nil {
		x.parent = y.parent
	}

	// 如果 y 没有父节点，则将 x 设为根节点
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	// 如果删除的是黑色节点，则需要进行平衡调整
	if y.color == black {
		t.deleteFixup(x, y.parent)
	}
}

// 删除后的红黑树性质的调整
func (t *redBlackTree) deleteFixup(x, px *node) {
	for x != t.root && (x == nil || x.color == black) {
		// 如果 x 是其父节点的左孩子
		if x == px.left {
			// 将 w 设为 x 的兄弟节点
			w := px.right
			// 如果 w 是红色的，情况 1
			if w.color == red {
				w.color = black
				px.color = red
				t.rotateLeft(px)
				w = px.right
			}

			// 如果 w 的两个子节点都是黑色的，情况 2
			if (w.left == nil || w.left.color == black) && (w.right == nil || w.right.color == black) {
				w.color = red
				x = px
				px = x.parent
			} else {
				// 如果 w 的右孩子是黑色的，情况 3
				if w.right == nil || w.right.color == black {
					w.left.color = black
					w.color = red
					t.rotateRight(w)
					w = px.right
				}

				// 情况 4
				w.color = px.color
				px.color = black
				w.right.color = black
				t.rotateLeft(px)
				x = t.root
			}
		} else { // 如果 x 是其父节点的右孩子
			// 将 w 设为 x 的兄弟节点
			w := px.left

			// 如果 w 是红色的，情况 1
			if w.color == red {
				w.color = black
				px.color = red
				t.rotateRight(px)
				w = px.left
			}

			// 如果 w 的两个子节点都是黑色的，情况 2
			if (w.left == nil || w.left.color == black) && (w.right == nil || w.right.color == black) {
				w.color = red
				x = px
				px = x.parent
			} else {
				// 如果 w 的左孩子是黑色的，情况 3
				if w.left == nil || w.left.color == black {
					w.right.color = black
					w.color = red
					t.rotateLeft(w)
					w = px.left
				}

				// 情况 4
				w.color = px.color
				px.color = black
				w.left.color = black
				t.rotateRight(px)
				x = t.root
			}
		}
	}

	if x != nil {
		x.color = black
	}
}

// 中序遍历
func (t *redBlackTree) inOrderTraverse(x *node) {
	if x != nil {
		t.inOrderTraverse(x.left)
		fmt.Print(x.val, " ")
		t.inOrderTraverse(x.right)
	}
}

func main() {
	// 创建一个红黑树
	t := newRedBlackTree()
	// 插入数据
	t.insert(11)
	t.insert(2)
	t.insert(14)
	t.insert(1)
	t.insert(7)
	t.insert(15)
	t.insert(5)
	t.insert(8)
	t.insert(4)
	t.insert(0)
	t.insert(9)
	t.insert(12)
	t.insert(13)
	t.insert(10)
	t.insert(6)
	t.insert(3)

	// 中序遍历
	t.inOrderTraverse(t.root)
	fmt.Println()

	// 删除数据
	t.delete(3)
	t.delete(5)
	t.delete(2)
	t.delete(1)
	t.delete(7)
	t.delete(4)

	// 中序遍历
	t.inOrderTraverse(t.root)
	fmt.Println()
}
