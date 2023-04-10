package main

import "fmt"

// 定义一个节点
type Node struct {
    data int    // 存储的数据
    prev *Node // 指向上一个节点的指针
    next *Node // 指向下一个节点的指针
}

// 定义一个双链表
type DoublyLinkedList struct {
    head *Node // 链表的头节点指针
}

// 初始化链表，设置链表的头节点
func (list *DoublyLinkedList) Init(data int) {
    list.head = &Node{data: data, prev: nil, next: nil}
}

// 向链表末尾添加新的节点
func (list *DoublyLinkedList) Add(data int) {
    newNode := &Node{data: data, prev: nil, next: nil}

    if list.head == nil { // 链表为空，直接将新节点作为头节点
        list.head = newNode
    } else { // 遍历链表，找到链表末尾，将新节点插入到末尾
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
        newNode.prev = current
    }
}

// 从链表中删除指定数据的节点
func (list *DoublyLinkedList) Remove(data int) {
    if list.head == nil { // 链表为空，直接返回
        return
    }

    if list.head.data == data { // 链表头节点就是要删除的节点
        list.head = list.head.next
        if list.head != nil {
            list.head.prev = nil
        }
    } else { // 遍历链表，找到要删除的节点，将其从链表中删除
        current := list.head
        for current != nil {
            if current.data == data {
                current.prev.next = current.next
                if current.next != nil {
                    current.next.prev = current.prev
                }
                return
            }
            current = current.next
        }
    }
}

// 在链表中查找指定数据的节点
func (list *DoublyLinkedList) Search(data int) bool {
    current := list.head
    for current != nil {
        if current.data == data {
            return true
        }
        current = current.next
    }
    return false
}

// 修改链表中的节点数据
func (list *DoublyLinkedList) Modify(oldData, newData int) {
    current := list.head
    for current != nil {
        if current.data == oldData {
            current.data = newData
            return
        }
        current = current.next
    }
}

// 打印链表中的所有节点数据
func (list *DoublyLinkedList) Print() {
    current := list.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    list := DoublyLinkedList{}
    list.Init(1) // 初始化链表，设置头节点

    list.Add(2)  // 添加新节点
    list.Add(3)
    list.Add(4)
    list.Add(5)
    list.Print()

    list.Remove(3) // 删除指定数据的节点
    list.Print()

    fmt.Println(list.Search(2)) // 查找指定数据的节点

    list.Modify(4, 6) // 修改节点的数据
    list.Print()
}
