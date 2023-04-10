package main

import "fmt"

func HeapSort(arr []int) {
    length := len(arr)
    // 从最后一个非叶子节点开始调整，构造最大堆
    for i := length/2 - 1; i >= 0; i-- {
        adjustHeap(arr, i, length)
    }
    // 不断将堆顶元素与最后一个元素交换，并重新调整堆，直到堆为空
    for i := length - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        adjustHeap(arr, 0, i)
    }
}

func adjustHeap(arr []int, i, length int) {
    temp := arr[i] // 当前节点的值
    for k := 2*i + 1; k < length; k = 2*k + 1 { // 从i节点的左子节点开始，找到最大的子节点
        if k+1 < length && arr[k] < arr[k+1] {
            k++
        }
        if arr[k] > temp { // 如果子节点大于当前节点，交换两个节点
            arr[i] = arr[k]
            i = k
        } else {
            break
        }
    }
    arr[i] = temp // 将当前节点放到合适的位置
}

func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Before sort:", arr)
    HeapSort(arr)
    fmt.Println("After sort:", arr)
}
