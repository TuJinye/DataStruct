package main

import "fmt"

func BinaryInsertSort(arr []int) []int {
    n := len(arr)
    if n <= 1 {
        return arr
    }
    for i := 1; i < n; i++ {
        temp := arr[i] // 用临时变量存储待排序的值
        low, high := 0, i-1
        for low <= high {
            mid := low + (high-low)/2 // 计算中间位置
            if arr[mid] > temp {
                high = mid - 1 // 在左半部分查找
            } else {
                low = mid + 1 // 在右半部分查找
            }
        }
        // 将mid右边的元素都向右移一位
        for j := i - 1; j >= low; j-- {
            arr[j+1] = arr[j]
        }
        arr[low] = temp // 插入待排序值
    }
    return arr
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    arr = BinaryInsertSort(arr)
    fmt.Println("After sort:", arr)
}
