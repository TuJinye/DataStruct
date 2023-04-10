package main

import "fmt"

func cycleSort(arr []int) {
    for cycleStart := 0; cycleStart < len(arr)-1; cycleStart++ {
        item := arr[cycleStart]
        pos := cycleStart
        for i := cycleStart + 1; i < len(arr); i++ {
            if arr[i] < item {
                pos++
            }
        }
        if pos == cycleStart {
            continue
        }
        for item == arr[pos] {
            pos++
        }
        arr[pos], item = item, arr[pos]
        for pos != cycleStart {
            pos = cycleStart
            for i := cycleStart + 1; i < len(arr); i++ {
                if arr[i] < item {
                    pos++
                }
            }
            for item == arr[pos] {
                pos++
            }
            arr[pos], item = item, arr[pos]
        }
    }
}

func main() {
    arr := []int{3, 4, 2, 1, 5}
    fmt.Println("Before sorting: ", arr)
    cycleSort(arr)
    fmt.Println("After sorting: ", arr)
}
