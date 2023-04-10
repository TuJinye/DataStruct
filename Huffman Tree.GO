package main

import (
	"container/heap"
	"fmt"
)

// 定义哈夫曼树的节点结构体
type Node struct {
	// 节点的值，对于叶子节点来说，这个值就是它代表的字符
	value rune
	// 节点的权重，也就是字符出现的次数
	weight int
	// 左子树和右子树
	left, right *Node
}

// 构造一个节点堆，实现heap.Interface接口
type NodeHeap []*Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push和Pop方法用于实现堆的插入和删除操作
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// 构造哈夫曼树
func buildHuffmanTree(frequencies map[rune]int) *Node {
	// 构造节点堆
	var nodes NodeHeap
	for c, f := range frequencies {
		nodes = append(nodes, &Node{c, f, nil, nil})
	}
	heap.Init(&nodes)

	// 不断取出堆顶的两个节点，合并成一个新节点，并重新插入堆中
	for nodes.Len() > 1 {
		left := heap.Pop(&nodes).(*Node)
		right := heap.Pop(&nodes).(*Node)
		parent := &Node{0, left.weight + right.weight, left, right}
		heap.Push(&nodes, parent)
	}

	// 返回根节点
	return heap.Pop(&nodes).(*Node)
}

// 递归遍历哈夫曼树，生成编码表
func buildCodeTable(root *Node, prefix []byte, codes map[rune][]byte) {
	if root == nil {
		return
	}

	if root.left == nil && root.right == nil {
		codes[root.value] = append([]byte{}, prefix...)
		return
	}

	prefix = append(prefix, '0')
	buildCodeTable(root.left, prefix, codes)
	prefix = prefix[:len(prefix)-1]

	prefix = append(prefix, '1')
	buildCodeTable(root.right, prefix, codes)
	prefix = prefix[:len(prefix)-1]
}

// 对一段文本进行编码
func encode(text string, codes map[rune][]byte) []byte {
	var result []byte
	for _, c := range text {
		result = append(result, codes[c]...)
	}
	return result
}

// 对一段编码后的文本进行解码
func decode(encoded []byte, root *Node) string {
	var result []r
	node := root
	for _, b := range encoded {
		if b == '0' {
			node = node.left
		} else {
			node = node.right
		}

		if node.left == nil && node.right == nil {
			result = append(result, node.value)
			node = root
		}
	}
	return string(result)
}

func main() {
	// 统计每个字符出现的次数
	frequencies := make(map[rune]int)
	text := "hello world"
	for _, c := range text {
		frequencies[c]++
	}
	// 构造哈夫曼树
	root := buildHuffmanTree(frequencies)

	// 构造编码表
	codes := make(map[rune][]byte)
	buildCodeTable(root, []byte{}, codes)

	// 对文本进行编码和解码
	encoded := encode(text, codes)
	decoded := decode(encoded, root)

	fmt.Println("Original text:", text)
	fmt.Println("Encoded text:", string(encoded))
	fmt.Println("Decoded text:", decoded)
}
