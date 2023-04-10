package main

import "fmt"

type Graph struct {
    V   int           // 图的顶点数
    Adj [][]int       // 存储邻接表
}

// 初始化图
func NewGraph(V int) *Graph {
    adj := make([][]int, V)
    for i := range adj {
        adj[i] = []int{}
    }
    return &Graph{V, adj}
}

// 添加边
func (g *Graph) AddEdge(v, w int) {
    g.Adj[v] = append(g.Adj[v], w)
}

// 深度优先遍历
func (g *Graph) DFS(start int) {
    visited := make([]bool, g.V)  // 用来记录节点是否已经被访问过
    g.dfs(start, visited)
}

// 递归遍历
func (g *Graph) dfs(v int, visited []bool) {
    visited[v] = true  // 标记该节点已经被访问过
    fmt.Printf("%d ", v)

    for _, w := range g.Adj[v] {
        if !visited[w] {
            g.dfs(w, visited)  // 递归遍历未被访问的邻接节点
        }
    }
}

func main() {
    g := NewGraph(6)
    g.AddEdge(0, 2)
    g.AddEdge(0, 1)
    g.AddEdge(1, 4)
    g.AddEdge(1, 3)
    g.AddEdge(2, 3)
    g.AddEdge(3, 5)
    g.AddEdge(3, 4)

    fmt.Println("深度优先遍历：")
    g.DFS(0)
}

