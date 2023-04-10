package main

import "fmt"

// 边节点
type EdgeNode struct {
	ivex  int       // 该边的终点位置
	ilink *EdgeNode // 指向下一条依附于该顶点的边
	jvex  int       // 该边的起点位置
	jlink *EdgeNode // 指向下一条依附于该顶点的边
}

// 顶点节点
type VertexNode struct {
	data      string    // 顶点信息
	firstedge *EdgeNode // 指向第一条依附于该顶点的边
}

// 邻接多重表
type AdjacencyMultiList struct {
	adjList     []VertexNode // 存放顶点的数组
	numVertexes int          // 顶点数
	numEdges    int          // 边数
}

// 创建邻接多重表
func NewAdjacencyMultiList(vertexes []string, edges [][2]int) *AdjacencyMultiList {
	// 获取顶点数和边数
	numVertexes := len(vertexes)
	numEdges := len(edges)

	// 创建邻接多重表
	adjList := make([]VertexNode, numVertexes)
	for i := 0; i < numVertexes; i++ {
		adjList[i].data = vertexes[i]
		adjList[i].firstedge = nil
	}
	for i := 0; i < numEdges; i++ {
		// 获取边的起点和终点
		ivex, jvex := edges[i][0], edges[i][1]

		// 创建边节点
		e := &EdgeNode{ivex: ivex, ilink: nil, jvex: jvex, jlink: nil}

		// 插入i到j的边
		if adjList[ivex].firstedge == nil {
			adjList[ivex].firstedge = e
		} else {
			// 找到i的邻接表的最后一个节点，将新的边节点插入到它后面
			p := adjList[ivex].firstedge
			for p.ilink != nil && p.ilink.ivex < ivex {
				p = p.ilink
			}
			e.ilink = p.ilink
			p.ilink = e
		}

		// 插入j到i的边
		if adjList[jvex].firstedge == nil {
			adjList[jvex].firstedge = e
		} else {
			// 找到j的邻接表的最后一个节点，将新的边节点插入到它后面
			p := adjList[jvex].firstedge
			for p.jlink != nil && p.jlink.jvex < jvex {
				p = p.jlink
			}
			e.jlink = p.jlink
			p.jlink = e
		}
	}

	return &AdjacencyMultiList{
		adjList:     adjList,
		numVertexes: numVertexes,
		numEdges:    numEdges,
	}
}

// 打印邻接多重表
func (g *AdjacencyMultiList) Print() {
	for i := 0; i < g.numVertexes; i++ {
		fmt.Printf("%d: %s ->", i, g.adjList[i].data)
		p := g.adjList[i].firstedge
		for p != nil {
			if p.ivex == i {
				fmt.Printf(" %d", p.jvex)
				p = p.ilink
			} else {
				fmt.Printf(" %d", p.ivex)
				p = p.jlink
			}
		}

		fmt.Println()
	}
}

// 测试邻接多重表
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
	g := NewAdjacencyMultiList(vertexes, edges)
	g.Print()
}
