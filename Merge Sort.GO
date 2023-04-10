package main

import "fmt"

func MergeSort(arr []int) {
    if len(arr) <= 1 { // 当序列长度为1时，停止递归
        return
    }
    mid := len(arr) / 2
    left := make([]int, mid) // 分别对左半部分和右半部分进行排序
    right := make([]int, len(arr)-mid)
    copy(left, arr[:mid])
    copy(right, arr[mid:])
    MergeSort(left)
    MergeSort(right)
    merge(arr, left, right) // 合并左半部分和右半部分
}

func merge(arr, left, right []int) {
    i, j, k := 0, 0, 0 // i, j分别指向左半部分和右半部分的第一个元素，k指向当前位置
    for i < len(left) && j < len(right) { // 将左半部分和右半部分中较小的元素依次放入arr中
        if left[i] <= right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }
    // 将左半部分或右半部分中剩余的元素依次放入arr中
    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }
    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    MergeSort(arr)
    fmt.Println("After sort:", arr)
}
