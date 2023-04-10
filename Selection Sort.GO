package main

import "fmt"

func SelectSort(arr []int) {
    length := len(arr)
    for i := 0; i < length-1; i++ {
        minIndex := i // 记录最小元素的下标
        for j := i + 1; j < length; j++ {
            if arr[j] < arr[minIndex] { // 找到更小的元素
                minIndex = j // 更新最小元素下标
            }
        }
        if minIndex != i { // 如果当前位置不是最小元素，交换两个元素
            arr[i], arr[minIndex] = arr[minIndex], arr[i]
        }
    }
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    SelectSort(arr)
    fmt.Println("After sort:", arr)
}
