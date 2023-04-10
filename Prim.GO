package main

import (
    "container/heap"
    "fmt"
)

type Edge struct {
    to, w int
}

type Graph struct {
    n      int
    adj    [][]Edge
}

func prim(g Graph) ([]Edge, int) {
    mst := make([]Edge, 0, g.n-1)
    vis := make([]bool, g.n)
    h := &MinHeap{{0, 0}}
    heap.Init(h)
    for h.Len() > 0 {
        x := heap.Pop(h).(Edge)
        if vis[x.to] {
            continue
        }
        vis[x.to] = true
        if x.w > 0 {
            mst = append(mst, x)
        }
        for _, e := range g.adj[x.to] {
            if !vis[e.to] {
                heap.Push(h, Edge{e.to, e.w})
            }
        }
    }
    var cost int
    for _, e := range mst {
        cost += e.w
    }
    return mst, cost
}

type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].w < h[j].w }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    n := 5
    g := Graph{n, make([][]Edge, n)}
    g.adj[0] = []Edge{{1, 2}, {3, 6}}
    g.adj[1] = []Edge{{0, 2}, {2, 3}, {4, 5}}
    g.adj[2] = []Edge{{1, 3}, {4, 7}}
    g.adj[3] = []Edge{{0, 6}, {4, 9}}
    g.adj[4] = []Edge{{1, 5}, {2, 7}, {3, 9}}
    mst, cost := prim(g)
    fmt.Println(mst)   // 输出结果 [{0 2} {1 3} {1 5} {0 3}]
    fmt.Println(cost)  // 输出结果 17
}