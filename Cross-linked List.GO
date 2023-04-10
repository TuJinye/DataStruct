package main

import "fmt"

type Node struct {
	row, col     int   // 节点对应的行列号
	right, down  *Node // 节点的右侧和下方指针
}

type CrossList struct {
	rows, cols   int      // 十字链表的行数和列数
	rowHeaders   []*Node  // 行头指针数组
	colHeaders   []*Node  // 列头指针数组
}

// 创建一个空的十字链表
func NewCrossList(rows, cols int) *CrossList {
	rowHeaders := make([]*Node, rows)
	colHeaders := make([]*Node, cols)
	for i := 0; i < rows; i++ {
		rowHeaders[i] = &Node{row: i}
	}
	for j := 0; j < cols; j++ {
		colHeaders[j] = &Node{col: j}
	}
	return &CrossList{rows: rows, cols: cols, rowHeaders: rowHeaders, colHeaders: colHeaders}
}

// 向十字链表中添加一个节点
func (cl *CrossList) AddNode(row, col int) {
	node := &Node{row: row, col: col}

	// 在对应行中添加节点
	prevRowNode := cl.rowHeaders[row]
	currRowNode := prevRowNode.right
	for currRowNode != nil && currRowNode.col < col {
		prevRowNode = currRowNode
		currRowNode = currRowNode.right
	}
	prevRowNode.right = node
	node.right = currRowNode

	// 在对应列中添加节点
	prevColNode := cl.colHeaders[col]
	currColNode := prevColNode.down
	for currColNode != nil && currColNode.row < row {
		prevColNode = currColNode
		currColNode = currColNode.down
	}
	prevColNode.down = node
	node.down = currColNode
}

// 输出十字链表
func (cl *CrossList) Print() {
	for i := 0; i < cl.rows; i++ {
		node := cl.rowHeaders[i].right
		for j := 0; j < cl.cols; j++ {
			if node != nil && node.col == j {
				fmt.Printf("%d ", 1)
				node = node.right
			} else {
				fmt.Printf("%d ", 0)
			}
		}
		fmt.Println()
	}
}

// 获取第row行中的节点
func (cl *CrossList) GetRow(row int) []*Node {
	nodes := make([]*Node, 0)
	node := cl.rowHeaders[row].right
	for node != nil {
		nodes = append(nodes, node)
		node = node.right
	}
	return nodes
}

// 获取第col列中的节点
func (cl *CrossList) GetCol(col int) []*Node {
	nodes := make([]*Node, 0)
	node := cl.colHeaders[col].down
	for node != nil {
		nodes = append(nodes, node)
		node = node.down
	}
	return nodes
}

func main() {
	cl := NewCrossList(5, 5)

	// 添加节点
	cl.AddNode(0, 1)
	cl.AddNode(0, 2)
	cl.AddNode(1, 0)
	cl.AddNode(2, 0)
	cl.AddNode(2, 1)
	// 输出十字链表
	cl.Print()

	// 获取第2行的节点
	nodes := cl.GetRow(2)
	fmt.Println("Nodes in row 2:", nodes)

	// 获取第1列的节点
	nodes = cl.GetCol(1)
	fmt.Println("Nodes in col 1:", nodes)
}