package main

import "fmt"

func InsertSort(arr []int) []int {
    n := len(arr)
    if n <= 1 {
        return arr
    }
    for i := 1; i < n; i++ {
        temp := arr[i] // 用临时变量存储待排序的值
        j := i - 1
        for j >= 0 && arr[j] > temp { // 将已排序数组中大于待排序值的数都向右移一位
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = temp // 插入待排序值
    }
    return arr
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    arr = InsertSort(arr)
    fmt.Println("After sort:", arr)
}
