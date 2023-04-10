package main

import "fmt"

func ShellSort(arr []int) []int {
    n := len(arr)
    if n <= 1 {
        return arr
    }
    gap := n / 2
    for gap > 0 {
        for i := gap; i < n; i++ {
            temp := arr[i] // 用临时变量存储待排序的值
            j := i - gap
            for j >= 0 && arr[j] > temp { // 将已排序数组中大于待排序值的数都向右移gap位
                arr[j+gap] = arr[j]
                j -= gap
            }
            arr[j+gap] = temp // 插入待排序值
        }
        gap /= 2 // 缩小增量
    }
    return arr
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    arr = ShellSort(arr)
    fmt.Println("After sort:", arr)
}
