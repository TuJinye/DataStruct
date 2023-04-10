// 定义二叉树节点结构体
type BinaryTreeNode struct {
    Value  int            // 节点值
    Left   *BinaryTreeNode // 左子节点
    Right  *BinaryTreeNode // 右子节点
}

// 插入节点
func (node *BinaryTreeNode) Insert(value int) {
    // 如果当前节点为空，则创建一个新节点
    if node == nil {
        node = &BinaryTreeNode{Value: value}
        return
    }
    // 如果插入值小于当前节点的值，则将它插入到左子树中
    if value < node.Value {
        if node.Left == nil {
            node.Left = &BinaryTreeNode{Value: value}
        } else {
            node.Left.Insert(value)
        }
    } else { // 否则将它插入到右子树中
        if node.Right == nil {
            node.Right = &BinaryTreeNode{Value: value}
        } else {
            node.Right.Insert(value)
        }
    }
}

// 查找节点
func (node *BinaryTreeNode) Search(value int) *BinaryTreeNode {
    if node == nil || node.Value == value { // 如果当前节点为空或者当前节点的值等于要查找的值，则返回当前节点
        return node
    }
    if value < node.Value { // 如果要查找的值小于当前节点的值，则在左子树中查找
        return node.Left.Search(value)
    } else { // 否则在右子树中查找
        return node.Right.Search(value)
    }
}

// 深度优先遍历（先序遍历）
func (node *BinaryTreeNode) TraverseDFS() {
    if node == nil {
        return
    }
    fmt.Printf("%d ", node.Value)
    node.Left.TraverseDFS()
    node.Right.TraverseDFS()
}

// 广度优先遍历（层序遍历）
func (node *BinaryTreeNode) TraverseBFS() {
    if node == nil {
        return
    }
    queue := []*BinaryTreeNode{node}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        fmt.Printf("%d ", current.Value)
        if current.Left != nil {
            queue = append(queue, current.Left)
        }
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
}
