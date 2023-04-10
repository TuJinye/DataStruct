package main

import (
    "fmt"
)

type LoserTree struct {
    K int
    Tree []int
}

func NewLoserTree(k int) *LoserTree {
    lt := &LoserTree{
        K: k,
        Tree: make([]int, k+1),
    }
    for i := range lt.Tree {
        lt.Tree[i] = -1
    }
    return lt
}

func (lt *LoserTree) Adjust(index int, input []int) {
    t := (index + lt.K) / 2
    for t > 0 {
        if index == -1 || (input[index] != -1 && input[index] < input[lt.Tree[t]]) {
            index, lt.Tree[t] = lt.Tree[t], index
        }
        t /= 2
    }
    lt.Tree[0] = index
}

func MergeSortedArrays(arrays [][]int) []int {
    k := len(arrays)
    lt := NewLoserTree(k)
    pointers := make([]int, k)
    for i := range pointers {
        pointers[i] = -1
    }
    for i := range arrays {
        if len(arrays[i]) > 0 {
            pointers[i] = 0
            lt.Adjust(i, arrays[i])
        }
    }
    var result []int
    for lt.Tree[0] != -1 {
        minIndex := lt.Tree[0]
        result = append(result, arrays[minIndex][pointers[minIndex]])
        pointers[minIndex]++
        if pointers[minIndex] >= len(arrays[minIndex]) {
            pointers[minIndex] = -1
            lt.Adjust(minIndex, []int{-1})
        } else {
            lt.Adjust(minIndex, arrays[minIndex])
        }
    }
    return result
}

func main() {
    arrays := [][]int{
        {3, 6, 9},
        {2, 4, 8, 10},
        {1, 5, 7},
    }
    merged := MergeSortedArrays(arrays)
    fmt.Println(merged)
}
