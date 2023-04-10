package main

import "fmt"

// 边的结构体
type EdgeNode struct {
	adjvex int       // 邻接点域，存储该顶点对应的下标
	weight int       // 权值
	next   *EdgeNode // 链域，指向下一个邻接点
}

// 顶点的结构体
type VertexNode struct {
	data      string    // 顶点域，存储顶点信息
	firstedge *EdgeNode // 边表头指针
}

// 邻接表结构体
type AdjacencyList struct {
	adjList     []*VertexNode // 邻接表
	numVertexes int           // 顶点数
	numEdges    int           // 边数
}

// 创建邻接表
func NewAdjacencyList(vertexes []string, edges [][2]int) *AdjacencyList {
	numVertexes := len(vertexes)
	adjList := make([]*VertexNode, numVertexes)

	// 初始化邻接表的顶点
	for i := 0; i < numVertexes; i++ {
		adjList[i] = &VertexNode{
			data:      vertexes[i],
			firstedge: nil,
		}
	}

	// 插入边
	for _, edge := range edges {
		i := edge[0]
		j := edge[1]

		// 插入i -> j的边
		e := &EdgeNode{
			adjvex: j,
			weight: 0,
			next:   adjList[i].firstedge,
		}
		adjList[i].firstedge = e

		// 插入j -> i的边
		e = &EdgeNode{
			adjvex: i,
			weight: 0,
			next:   adjList[j].firstedge,
		}
		adjList[j].firstedge = e
	}

	return &AdjacencyList{
		adjList:     adjList,
		numVertexes: numVertexes,
		numEdges:    len(edges),
	}
}

// 打印邻接表
func (g *AdjacencyList) Print() {
	for i := 0; i < g.numVertexes; i++ {
		fmt.Printf("%d: %s ->", i, g.adjList[i].data)

		p := g.adjList[i].firstedge
		for p != nil {
			fmt.Printf(" %d", p.adjvex)
			p = p.next
		}

		fmt.Println()
	}
}

// 测试邻接表
func main() {
	vertexes := []string{"V0", "V1", "V2", "V3", "V4", "V5"}
	edges := [][2]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{3, 4},
		{3, 5},
	}

	g := NewAdjacencyList(vertexes, edges)
	g.Print()
}
