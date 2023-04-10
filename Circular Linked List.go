package main

import "fmt"

// 循环链表节点
type Node struct {
	data interface{}
	next *Node
}

// 循环链表
type CircularLinkedList struct {
	head   *Node
	length int
}

// 在链表末尾添加节点
func (c *CircularLinkedList) Append(data interface{}) {
	// 创建一个新节点
	newNode := &Node{data, nil}

	if c.head == nil {
		// 如果链表为空，则将新节点作为头部节点，并使其 next 指向自身
		c.head = newNode
		newNode.next = newNode
	} else {
		// 否则，从头部节点开始遍历链表，找到末尾节点，并使其 next 指向新节点
		current := c.head
		for current.next != c.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = c.head
	}

	// 增加链表的长度
	c.length++
}

// 在链表开头添加节点
func (c *CircularLinkedList) Prepend(data interface{}) {
	// 创建一个新节点
	newNode := &Node{data, nil}

	if c.head == nil {
		// 如果链表为空，则将新节点作为头部节点，并使其 next 指向自身
		c.head = newNode
		newNode.next = newNode
	} else {
		// 否则，将新节点作为头部节点，并使其 next 指向原来的头部节点
		current := c.head
		for current.next != c.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = c.head
		c.head = newNode
	}

	// 增加链表的长度
	c.length++
}

// 在指定位置插入节点
func (c *CircularLinkedList) Insert(data interface{}, index int) {
	// 如果链表为空或者索引无效，则不执行任何操作
	if c.head == nil ||
		index < 0 || index > c.length {
		return
	}

	// 如果插入位置为 0，则在链表开头插入节点
	if index == 0 {
		c.Prepend(data)
		return
	}

	// 如果插入位置为链表末尾，则在链表末尾插入节点
	if index == c.length {
		c.Append(data)
		return
	}

	// 创建一个新节点
	newNode := &Node{data, nil}

	// 遍历链表，找到指定位置的节点和其前一个节点
	current := c.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	prev := current
	current = current.next

	// 将新节点的 next 指针指向指定位置的节点，将其前一个节点的 next 指针指向新节点
	newNode.next = current
	prev.next = newNode

	// 增加链表的长度
	c.length++
}

// 删除指定位置的节点
func (c *CircularLinkedList) Remove(index int) {
	// 如果链表为空或者索引无效，则不执行任何操作
	if c.head == nil ||
		index < 0 || index >= c.length {
		return
	}
	// 如果删除的是头部节点，则将头部指针指向下一个节点
	if index == 0 {
		current := c.head
		for current.next != c.head {
			current = current.next
		}
		c.head = c.head.next
		current.next = c.head
	} else {
		// 否则，遍历链表，找到指定位置的节点和其前一个节点，并将前一个节点的 next 指针指向指定位置的下一个节点
		current := c.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		prev := current
		current = current.next
		prev.next = current.next
	}

	// 减少链表的长度
	c.length--
}

// 获取链表长度
func (c *CircularLinkedList) Len() int {
	return c.length
}

// 打印链表中所有节点的值
func (c *CircularLinkedList) Print() {
	if c.head == nil {
		fmt.Println("链表为空")
	} else {
		current := c.head
		for current.next != c.head {
			fmt.Printf("%v -> ", current.data)
			current = current.next
		}
		fmt.Printf("%v -> ...\n", current.data)
	}
}

// 示例
func main() {
	// 创建一个新的循环链表
	cl := &CircularLinkedList{}
	// 在链表末尾添加节点
	cl.Append(1)
	cl.Append(2)
	cl.Append(3)
	cl.Append(4)

	// 在链表开头添加节点
	cl.Prepend(0)

	// 在指定位置插入节点
	cl.Insert(5, 5)

	// 打印链表中所有节点的值
	cl.Print()

	// 删除指定位置的节点
	cl.Remove(4)

	// 获取链表长度
	fmt.Println("链表长度为：", cl.Len())

	// 打印链表中所有节点的值
	cl.Print()
}
