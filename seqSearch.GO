package main

import "fmt"

func seqSearch(arr []int, key int) int {
    for i := 0; i < len(arr); i++ {
        if arr[i] == key {
            return i // 返回找到的元素的下标
        }
    }
    return -1 // 没有找到，返回-1
}

func main() {
    arr := []int{5, 2, 4, 6, 1, 3}
    key := 4
    index := seqSearch(arr, key)
    if index == -1 {
        fmt.Println("未找到元素", key)
    } else {
        fmt.Println("元素", key, "的下标为", index)
    }
}
