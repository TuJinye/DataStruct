package main

import "fmt"

// 折半查找
func BinarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1  // 初始化左右指针
    for left <= right {  // 只要左指针不大于右指针就继续查找
        mid := (left + right) / 2  // 中间位置
        if arr[mid] == target {
            return mid  // 如果中间位置元素等于目标值，返回中间位置下标
        } else if arr[mid] > target {
            right = mid - 1  // 如果中间位置元素大于目标值，更新右指针
        } else {
            left = mid + 1  // 如果中间位置元素小于目标值，更新左指针
        }
    }
    return -1  // 如果未找到目标值，返回-1
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11}
    target := 7
    index := BinarySearch(arr, target)
    if index == -1 {
        fmt.Println("未找到目标值")
    } else {
        fmt.Printf("目标值 %d 在数组中的下标为 %d\n", target, index)
    }
}

