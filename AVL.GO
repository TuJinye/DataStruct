package avltree

// avl树的节点结构体
type node struct {
	value  int   // 节点的值
	height int   // 子树高度
	left   *node // 左子节点
	right  *node // 右子节点
}

// AVL树结构体
type AVLTree struct {
	root *node // 根节点
}

// 返回节点的高度
func height(n *node) int {
	if n == nil {
		return -1
	}
	return n.height
}

// 计算左右子树高度差
func balanceFactor(n *node) int {
	return height(n.left) - height(n.right)
}

// 左旋转
func rotateLeft(n *node) *node {
	r := n.right
	n.right = r.left
	r.left = n

	// 更新节点高度
	n.height = max(height(n.left), height(n.right)) + 1
	r.height = max(height(r.left), height(r.right)) + 1

	return r
}

// 右旋转
func rotateRight(n *node) *node {
	l := n.left
	n.left = l.right
	l.right = n

	// 更新节点高度
	n.height = max(height(n.left), height(n.right)) + 1
	l.height = max(height(l.left), height(l.right)) + 1

	return l
}

// 右左双旋转
func rotateRightLeft(n *node) *node {
	n.right = rotateRight(n.right)
	return rotateLeft(n)
}

// 左右双旋转
func rotateLeftRight(n *node) *node {
	n.left = rotateLeft(n.left)
	return rotateRight(n)
}

// 插入一个值到AVL树中
func (t *AVLTree) Insert(value int) {
	t.root = insertNode(t.root, value)
}

// 插入一个节点到AVL树中
func insertNode(n *node, value int) *node {
	// 如果为空节点，则直接返回一个新节点
	if n == nil {
		return &node{value, 0, nil, nil}
	}

	// 如果插入的值小于当前节点的值，往左子树插入
	if value < n.value {
		n.left = insertNode(n.left, value)

		// 检查是否需要旋转平衡
		if balanceFactor(n) == 2 {
			if balanceFactor(n.left) == 1 {
				n = rotateRight(n)
			} else {
				n = rotateLeftRight(n)
			}
		}
	} else if value > n.value { // 如果插入的值大于当前节点的值，往右子树插入
		n.right = insertNode(n.right, value)

		// 检查是否需要旋转平衡
		if balanceFactor(n) == -2 {
			if balanceFactor(n.right) == -1 {
				n = rotateLeft(n)
			} else {
				n = rotateRightLeft(n)
			}
		}
	}

	// 更新节点高度
	n.height = max(height(n.left), height(n.right)) + 1

	return n
}

// 删除节点
func (t *AVLTree) delete(node *AVLNode, key int) *AVLNode {
    if node == nil {
        return nil
    }
    if key < node.key {
        node.left = t.delete(node.left, key)
    } else if key > node.key {
        node.right = t.delete(node.right, key)
    } else { // 找到要删除的节点
        if node.left == nil && node.right == nil { // 如果是叶子节点，直接删除
            node = nil
        } else if node.left == nil { // 如果只有右子节点，用右子节点替换该节点
            node = node.right
        } else if node.right == nil { // 如果只有左子节点，用左子节点替换该节点
            node = node.left
        } else { // 如果左右子节点都有，用中序遍历的前驱或后继节点来替换该节点
            // 选用前驱节点
            pre := node.left
            for pre.right != nil {
                pre = pre.right
            }
            node.key, node.value = pre.key, pre.value // 用前驱节点的值来替换该节点的值
            node.left = t.delete(node.left, pre.key) // 在左子树中删除前驱节点
        }
    }

    // 如果删除节点后树为空，则直接返回
    if node == nil {
        return nil
    }

    // 更新节点高度
    node.height = max(t.getHeight(node.left), t.getHeight(node.right)) + 1

    // 计算节点平衡因子
    balanceFactor := t.getBalanceFactor(node)

    // 平衡处理
    // 1. LL 左左型，进行右旋
    if balanceFactor > 1 && t.getBalanceFactor(node.left) >= 0 {
        return t.rightRotate(node)
    }

    // 2. LR 左右型，先进行左旋，再进行右旋
    if balanceFactor > 1 && t.getBalanceFactor(node.left) < 0 {
        node.left = t.leftRotate(node.left)
        return t.rightRotate(node)
    }

    // 3. RR 右右型，进行左旋
    if balanceFactor < -1 && t.getBalanceFactor(node.right) <= 0 {
        return t.leftRotate(node)
    }

    // 4. RL 右左型，先进行右旋，再进行左旋
    if balanceFactor < -1 && t.getBalanceFactor(node.right) > 0 {
        node.right = t.rightRotate(node.right)
        return t.leftRotate(node)
    }

    return node
}

