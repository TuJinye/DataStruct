package main

import (
	"fmt"
)

// 定义一个节点结构体
type Node struct {
	value int     // 节点的值
	child []*Node // 子节点的指针数组
}

// 定义森林结构体
type Forest []*Node

// 定义森林的插入操作
func (f *Forest) Insert(parent int, value int) {
	// 如果父节点编号等于-1，则新建一个树，插入一个节点
	if parent == -1 {
		newTree := &Node{value: value}
		*f = append(*f, newTree)
	} else {
		// 在父节点的子节点中插入一个新节点
		newNode := &Node{value: value}
		(*f)[parent].child = append((*f)[parent].child, newNode)
	}
}

// 定义森林的遍历操作
func (f *Forest) Traverse() {
	// 遍历每一个树
	for i := 0; i < len(*f); i++ {
		fmt.Print("Tree ", i, ": ")
		// 遍历树中的每个节点
		f.traverseNode((*f)[i])
		fmt.Println()
	}
}

// 遍历树中的节点
func (f *Forest) traverseNode(n *Node) {
	// 输出节点的值
	fmt.Print(n.value, " ")
	// 遍历子节点
	for i := 0; i < len(n.child); i++ {
		f.traverseNode(n.child[i])
	}
}

func main() {
	// 创建一个森林
	f := Forest{}
	// 插入节点
	f.Insert(-1, 1)  // 在树0中插入节点1
	f.Insert(0, 2)   // 在树0的节点1中插入节点2
	f.Insert(0, 3)   // 在树0的节点1中插入节点3
	f.Insert(-1, 4)  // 在树1中插入节点4
	f.Insert(1, 5)   // 在树0的节点2中插入节点5
	f.Insert(1, 6)   // 在树0的节点2中插入节点6
	f.Insert(2, 7)   // 在树0的节点3中插入节点7
	f.Insert(3, 8)   // 在树1的节点1中插入节点8
	// 遍历森林
	f.Traverse()
}
