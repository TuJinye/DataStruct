package hashtable

// Node 表示哈希表中的一个节点
type Node struct {
	Key   int
	Value interface{}
	Next  *Node
}

// HashTable 表示哈希表
type HashTable struct {
	Size int
	Data []*Node
}

// New 创建一个指定大小的哈希表
func New(size int) *HashTable {
	return &HashTable{
		Size: size,
		Data: make([]*Node, size),
	}
}

// hash 将key哈希为哈希表中的位置
func (ht *HashTable) hash(key int) int {
	return key % ht.Size
}

// Put 往哈希表中插入一条记录
func (ht *HashTable) Put(key int, value interface{}) {
	hash := ht.hash(key)
	node := &Node{Key: key, Value: value}

	if ht.Data[hash] == nil {
		ht.Data[hash] = node
	} else {
		// 链式解决冲突
		cur := ht.Data[hash]
		for cur.Next != nil {
			if cur.Key == key {
				cur.Value = value
				return
			}
			cur = cur.Next
		}
		if cur.Key == key {
			cur.Value = value
			return
		}
		cur.Next = node
	}
}

// Get 从哈希表中获取对应key的记录的值
func (ht *HashTable) Get(key int) interface{} {
	hash := ht.hash(key)
	if ht.Data[hash] == nil {
		return nil
	} else {
		// 遍历链表查找记录
		cur := ht.Data[hash]
		for cur != nil {
			if cur.Key == key {
				return cur.Value
			}
			cur = cur.Next
		}
		return nil
	}
}

// Remove 从哈希表中删除对应key的记录
func (ht *HashTable) Remove(key int) {
	hash := ht.hash(key)
	if ht.Data[hash] == nil {
		return
	} else {
		// 遍历链表删除记录
		var prev *Node
		cur := ht.Data[hash]
		for cur != nil {
			if cur.Key == key {
				if prev == nil {
					ht.Data[hash] = cur.Next
				} else {
					prev.Next = cur.Next
				}
				return
			}
			prev = cur
			cur = cur.Next
		}
	}
}
