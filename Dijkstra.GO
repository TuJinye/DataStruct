package main

import (
    "container/heap"  // 用于实现优先队列
    "fmt"
)

const INF = 1e9  // 定义正无穷

type Edge struct {
    to, cost int
}

type State struct {
    distance int
    node     int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(State)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func dijkstra(graph [][]Edge, start int) []int {
    n := len(graph)
    dist := make([]int, n)
    for i := range dist {
        dist[i] = INF  // 初始化距离为正无穷
    }
    dist[start] = 0

    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, State{0, start})

    for pq.Len() > 0 {
        s := heap.Pop(pq).(State)
        u := s.node

        for _, e := range graph[u] {
            v := e.to
            cost := e.cost
            if dist[v] > dist[u]+cost {
                dist[v] = dist[u] + cost
                heap.Push(pq, State{dist[v], v})
            }
        }
    }

    return dist
}

func main() {
    // 构建图
    graph := make([][]Edge, 4)
    graph[0] = []Edge{{1, 1}, {2, 4}}
    graph[1] = []Edge{{2, 2}, {3, 5}}
    graph[2] = []Edge{{3, 1}}
    graph[3] = []Edge{}

    start := 0
    dist := dijkstra(graph, start)
    fmt.Println(dist)  // 输出结果 [0 1 3 4]
}
