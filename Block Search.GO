package main

import (
    "fmt"
    "math"
    "sort"
)

// 分块查找
func BlockSearch(arr []int, blockSize int, target int) int {
    blockNum := int(math.Ceil(float64(len(arr)) / float64(blockSize)))  // 计算块的数量
    blockMaxs := make([]int, blockNum)  // 存储每个块的最大值
    for i := 0; i < blockNum; i++ {
        blockStart := i * blockSize
        blockEnd := (i+1)*blockSize - 1
        if blockEnd >= len(arr) {
            blockEnd = len(arr) - 1
        }
        blockMaxs[i] = arr[blockEnd]  // 将每个块的最大值存储到blockMaxs数组中
    }
    blockIndex := sort.SearchInts(blockMaxs, target)  // 使用标准库的二分查找函数查找目标值所在的块
    if blockIndex == blockNum || arr[blockIndex*blockSize] > target {  // 目标值比第一个元素小或大于最后一个块的最大值，则未找到
        return -1
    }
    blockStart := blockIndex * blockSize
    blockEnd := (blockIndex+1)*blockSize - 1
    if blockEnd >= len(arr) {
        blockEnd = len(arr) - 1
    }
    for i := blockStart; i <= blockEnd; i++ {  // 在目标块中使用顺序查找寻找目标值
        if arr[i] == target {
            return i
        }
    }
    return -1  // 如果未找到目标值，返回-1
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    blockSize := 3
    target := 7
    index := BlockSearch(arr, blockSize, target)
    if index == -1 {
        fmt.Println("未找到目标值")
    } else {
        fmt.Printf("目标值 %d 在数组中的下标为 %d\n", target, index)
    }
}

