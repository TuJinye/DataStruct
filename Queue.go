package main

import (
    "fmt"
)

// 定义队列结构体
type Queue struct {
    items []int
}

// 入队
func (q *Queue) Enqueue(i int) {
    q.items = append(q.items, i)
}

// 出队
func (q *Queue) Dequeue() int {
    toRemove := q.items[0]    // 获取队列头部元素
    q.items = q.items[1:]     // 移除队列头部元素
    return toRemove
}

// 返回队列的大小
func (q *Queue) Size() int {
    return len(q.items)
}

// 返回队列是否为空
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

// 返回队列的头部元素
func (q *Queue) Peek() int {
    return q.items[0]
}

// 清空队列
func (q *Queue) Clear() {
    q.items = []int{}
}

func main() {
    // 创建一个新的队列
    q := Queue{}

    // 将元素添加到队列中
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    // 输出队列头部元素
    fmt.Println(q.Peek())

    // 输出队列大小
    fmt.Println(q.Size())

    // 从队列中移除元素
    q.Dequeue()

    // 输出队列是否为空
    fmt.Println(q.IsEmpty())

    // 清空队列
    q.Clear()

    // 输出队列大小
    fmt.Println(q.Size())
}
