// 定义树节点结构体
type TreeNode struct {
    Value    int         // 节点值
    Children []*TreeNode // 子节点切片
}

// 在树上添加子节点
func (node *TreeNode) AddChild(child *TreeNode) {
    node.Children = append(node.Children, child)
}

// 在树上删除子节点
func (node *TreeNode) RemoveChild(child *TreeNode) {
    for i, n := range node.Children {
        if n == child {
            node.Children = append(node.Children[:i], node.Children[i+1:]...)
            return
        }
    }
}

// 深度优先遍历（先序遍历）
func (node *TreeNode) TraverseDFS() {
    fmt.Printf("%d ", node.Value)
    for _, child := range node.Children {
        child.TraverseDFS()
    }
}

// 广度优先遍历（层序遍历）
func (node *TreeNode) TraverseBFS() {
    queue := []*TreeNode{node}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        fmt.Printf("%d ", current.Value)
        for _, child := range current.Children {
            queue = append(queue, child)
        }
    }
}
