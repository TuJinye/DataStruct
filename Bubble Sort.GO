package main

import "fmt"

func BubbleSort(arr []int) []int {
    n := len(arr)
    if n <= 1 {
        return arr
    }
    for i := 0; i < n-1; i++ {
        flag := true // 标记，如果一趟下来没有发生交换，则表示已经排好序了，退出循环
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] // 交换两个元素
                flag = false
            }
        }
        if flag {
            break
        }
    }
    return arr
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    arr = BubbleSort(arr)
    fmt.Println("After sort:", arr)
}
