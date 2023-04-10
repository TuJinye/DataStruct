package main

import "fmt"

func QuickSort(arr []int, left, right int) {
    if left < right {
        i, j := left, right
        pivot := arr[(left+right)/2] // 中间的数作为基准值
        for {
            for arr[i] < pivot { // 找到左边比基准值大的数
                i++
            }
            for arr[j] > pivot { // 找到右边比基准值小的数
                j--
            }
            if i >= j { // 退出循环条件
                break
            }
            arr[i], arr[j] = arr[j], arr[i] // 交换两个数
            i++
            j--
        }
        QuickSort(arr, left, i-1)  // 对左半部分进行快速排序
        QuickSort(arr, j+1, right) // 对右半部分进行快速排序
    }
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    QuickSort(arr, 0, len(arr)-1)
    fmt.Println("After sort:", arr)
}
