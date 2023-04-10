package main

import (
    "fmt"
    "sort"
)

type Edge struct {
    u, v, w int
}

func kruskal(n int, edges []Edge) ([]Edge, int) {
    uf := NewUnionFind(n)
    sort.Slice(edges, func(i, j int) bool {
        return edges[i].w < edges[j].w
    })
    var mst []Edge
    var cost int
    for _, e := range edges {
        if uf.Find(e.u) != uf.Find(e.v) { // 判断加入该边是否会形成环路
            uf.Union(e.u, e.v)
            mst = append(mst, e)
            cost += e.w
        }
    }
    return mst, cost
}

type UnionFind struct {
    parent, rank []int
}

func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{make([]int, n), make([]int, n)}
    for i := range uf.parent {
        uf.parent[i] = i  // 初始化每个节点的根节点为它自己
    }
    return uf
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])  // 路径压缩
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX == rootY {
        return
    }
    if uf.rank[rootX] < uf.rank[rootY] {
        rootX, rootY = rootY, rootX
    }
    uf.parent[rootY] = rootX
    if uf.rank[rootX] == uf.rank[rootY] {
        uf.rank[rootX]++
    }
}

func main() {
    n := 5
    edges := []Edge{{0, 1, 2}, {0, 3, 6}, {1, 3, 8}, {1, 2, 3}, {1, 4, 5}, {2, 4, 7}, {3, 4, 9}}
    mst, cost := kruskal(n, edges)
    fmt.Println(mst)   // 输出结果 [{1 2 3} {0 1 2} {1 4 5} {0 3 6}]
    fmt.Println(cost)  // 输出结果 16
}
