package main

import "fmt"

type Graph struct {
	V   int      // 图的顶点数
	E   int      // 图的边数
	adj [][]bool // 邻接矩阵
}

// 创建一个有V个顶点但没有边的图
func NewGraph(V int) *Graph {
	g := &Graph{V: V, adj: make([][]bool, V)}
	for i := 0; i < V; i++ {
		g.adj[i] = make([]bool, V)
	}
	return g
}

// 向图中添加一条边v-w
func (g *Graph) AddEdge(v int, w int) {
	g.adj[v][w] = true
	g.adj[w][v] = true // 无向图需要将两个方向都设置
	g.E++
}

// 判断图中是否存在边v-w
func (g *Graph) HasEdge(v int, w int) bool {
	return g.adj[v][w]
}

// 返回节点v的所有相邻节点
func (g *Graph) Adj(v int) []int {
	adj := make([]int, 0)
	for i := 0; i < g.V; i++ {
		if g.adj[v][i] {
			adj = append(adj, i)
		}
	}
	return adj
}

func main() {
	g := NewGraph(5)

	// 添加边
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)

	// 输出邻接矩阵
	fmt.Println("邻接矩阵：")
	for i := 0; i < g.V; i++ {
		for j := 0; j < g.V; j++ {
			if g.adj[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}

	// 输出节点的相邻节点
	fmt.Println("节点的相邻节点：")
	for i := 0; i < g.V; i++ {
		fmt.Printf("%d: ", i)
		adj := g.Adj(i)
		for _, v := range adj {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
