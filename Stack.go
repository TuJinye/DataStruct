package main

import (
    "fmt"
)

// 定义栈结构体
type Stack struct {
    items []int
}

// 入栈操作
func (s *Stack) Push(i int) {
    s.items = append(s.items, i)
}

// 出栈操作
func (s *Stack) Pop() int {
    // 判断栈是否为空
    if len(s.items) == 0 {
        fmt.Println("Stack is empty!")
        return -1
    }
    // 获取栈顶元素并移除
    popItem := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return popItem
}

// 获取栈顶元素
func (s *Stack) Peek() int {
    // 判断栈是否为空
    if len(s.items) == 0 {
        fmt.Println("Stack is empty!")
        return -1
    }
    return s.items[len(s.items)-1]
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

// 获取栈的大小
func (s *Stack) Size() int {
    return len(s.items)
}

func main() {
    // 创建一个新栈
    myStack := Stack{}

    // 入栈操作
    myStack.Push(10)
    myStack.Push(20)
    myStack.Push(30)

    // 输出栈顶元素
    fmt.Println("Peek:", myStack.Peek())

    // 出栈操作
    fmt.Println("Pop:", myStack.Pop())

    // 输出栈的大小
    fmt.Println("Size:", myStack.Size())

    // 判断栈是否为空
    fmt.Println("Is empty:", myStack.IsEmpty())
}
