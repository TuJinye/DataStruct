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

// 广度优先遍历
func (g *Graph) BFS(start int) {
    visited := make([]bool, g.V)  // 用来记录节点是否已经被访问过
    queue := []int{}              // 用来存储待遍历的节点
    visited[start] = true         // 将起始节点标记为已访问
    queue = append(queue, start)  // 将起始节点添加到待遍历队列中

    for len(queue) > 0 {
        // 取出队列头部节点并打印
        v := queue[0]
        fmt.Printf("%d ", v)
        queue = queue[1:]

        // 遍历该节点的邻接节点，将未被访问的节点添加到队列中
        for _, w := range g.Adj[v] {
            if !visited[w] {
                visited[w] = true
                queue = append(queue, w)
            }
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

    fmt.Println("广度优先遍历：")
    g.BFS(0)
}

