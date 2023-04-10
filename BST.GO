package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 在BST中插入节点
func Insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = Insert(root.Left, val)
    } else if val > root.Val {
        root.Right = Insert(root.Right, val)
    }
    return root
}

// 在BST中删除节点
func Delete(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if val < root.Val {
        root.Left = Delete(root.Left, val)
    } else if val > root.Val {
        root.Right = Delete(root.Right, val)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        minNode := root.Right
        for minNode.Left != nil {
            minNode = minNode.Left
        }
        root.Val = minNode.Val
        root.Right = Delete(root.Right, minNode.Val)
    }
    return root
}

// 在BST中查找节点
func Search(root *TreeNode, val int) bool {
    if root == nil {
        return false
    }
    if root.Val == val {
        return true
    } else if val < root.Val {
        return Search(root.Left, val)
    } else {
        return Search(root.Right, val)
    }
}

// 中序遍历BST
func InOrderTraversal(root *TreeNode) []int {
    res := []int{}
    if root != nil {
        res = append(res, InOrderTraversal(root.Left)...)
        res = append(res, root.Val)
        res = append(res, InOrderTraversal(root.Right)...)
    }
    return res
}

func main() {
    root := &TreeNode{Val: 5}
    Insert(root, 3)
    Insert(root, 2)
    Insert(root, 4)
    Insert(root, 7)
    Insert(root, 6)
    Insert(root, 8)
    fmt.Println("BST 中序遍历：", InOrderTraversal(root))  // 输出：[2 3 4 5 6 7 8]
    fmt.Println("查找节点 6 是否存在：", Search(root, 6))  // 输出：true
    root = Delete(root, 5)
    fmt.Println("删除节点 5 后 BST 中序遍历：", InOrderTraversal(root))  // 输出：[2 3 4 6 7 8]
}
