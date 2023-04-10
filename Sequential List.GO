package main

import "fmt"

const maxListSize = 100 // 最大列表长度

type SeqList struct {
    data [maxListSize]int // 存储元素的数组
    length int            // 当前元素个数
}

// 初始化一个空列表
func InitList() *SeqList {
    return &SeqList{length: 0}
}

// 获取列表长度
func (list *SeqList) Length() int {
    return list.length
}

// 获取指定位置的元素
func (list *SeqList) Get(index int) (int, error) {
    if index < 0 || index >= list.length {
        return 0, fmt.Errorf("index out of range")
    }
    return list.data[index], nil
}

// 在指定位置插入元素
func (list *SeqList) Insert(index int, value int) error {
    if index < 0 || index > list.length {
        return fmt.Errorf("index out of range")
    }
    if list.length >= maxListSize {
        return fmt.Errorf("list is full")
    }
    for i := list.length; i > index; i-- {
        list.data[i] = list.data[i-1]
    }
    list.data[index] = value
    list.length++
    return nil
}

// 删除指定位置的元素
func (list *SeqList) Delete(index int) error {
    if index < 0 || index >= list.length {
        return fmt.Errorf("index out of range")
    }
    for i := index; i < list.length-1; i++ {
        list.data[i] = list.data[i+1]
    }
    list.length--
    return nil
}

func main() {
    list := InitList()
    list.Insert(0, 1)
    list.Insert(1, 2)
    list.Insert(2, 3)
    fmt.Println(list.Length()) // 输出：3
    fmt.Println(list.Get(1))   // 输出：2
    list.Delete(1)
    fmt.Println(list.Length()) // 输出：2
    fmt.Println(list.Get(1))   // 输出：3
}
