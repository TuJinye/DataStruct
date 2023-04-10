package main

import (
    "fmt"
)

func RadixSort(arr []int) {
    if len(arr) <= 1 { // 当序列长度为1时，停止递归
        return
    }
    maxVal := getMaxVal(arr) // 获取序列中的最大值
    for digit := 1; maxVal/digit > 0; digit *= 10 { // 对每一位进行排序，直到最高位
        countingSort(arr, digit) // 使用计数排序进行排序
    }
}

func getMaxVal(arr []int) int {
    maxVal := arr[0]
    for _, val := range arr[1:] { // 遍历序列，找到最大值
        if val > maxVal {
            maxVal = val
        }
    }
    return maxVal
}

func countingSort(arr []int, digit int) {
    n := len(arr)
    output := make([]int, n)
    count := make([]int, 10) // 创建计数数组，记录当前位每个数字出现的次数
    for _, val := range arr {
        index := (val / digit) % 10
        count[index]++
    }
    for i := 1; i < 10; i++ { // 计算每个数字在排序后的序列中的位置
        count[i] += count[i-1]
    }
    for i := n - 1; i >= 0; i-- { // 从后往前遍历序列，放入output数组中
        index := (arr[i] / digit) % 10
        output[count[index]-1] = arr[i]
        count[index]--
    }
    for i := 0; i < n; i++ { // 将排序好的序列放回arr中
        arr[i] = output[i]
    }
}

func main() {
    arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
    fmt.Println("Before sort:", arr)
    RadixSort(arr)
    fmt.Println("After sort:", arr)
}
