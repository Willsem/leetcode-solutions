type listNode struct {
	value int
	key   int

	next *listNode
	prev *listNode
}

type LRUCache struct {
	capacity int
	values   map[int]*listNode

	// tail -> ... -> head
	// head first to delete
	head *listNode
	tail *listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		values:   make(map[int]*listNode, 0),
	}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.values[key]; !ok {
		return -1
	}

	node := c.values[key]
	c.putToStartQueue(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.values[key]; ok {
		c.putToStartQueue(node)
		node.value = value
		return
	}

	node := &listNode{
		value: value,
		key:   key,
	}
	c.values[key] = node

	if len(c.values)-1 == 0 {
		c.head = node
		c.tail = node
		return
	}

	if c.capacity == 1 {
		temp := c.head
		c.head = node
		c.tail = node
		delete(c.values, temp.key)
		return
	}

	if len(c.values)-1 >= c.capacity {
		temp := c.head
		c.head = c.head.prev
		c.head.next = nil

		delete(c.values, temp.key)
	}

	c.tail.prev = node
	node.next = c.tail
	c.tail = node
}

func (c *LRUCache) putToStartQueue(node *listNode) {
	if c.head == c.tail && c.head == node || node == c.tail {
		return
	}

	if c.head == node {
		c.head = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	c.tail.prev = node
	node.next = c.tail
	node.prev = nil
	c.tail = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
