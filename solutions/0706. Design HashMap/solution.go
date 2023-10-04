type node struct {
	value int
	key   int
	next  *node
}

type MyHashMap struct {
	data []*node
	size int
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]*node, 1000),
		size: 1000,
	}
}

func (m *MyHashMap) Put(key int, value int) {
	calcKey := m.hash(key)

	if m.data[calcKey] == nil {
		m.data[calcKey] = &node{
			value: value,
			key:   key,
		}
		return
	}

	var prev *node
	temp := m.data[calcKey]
	for temp != nil && temp.key != key {
		prev = temp
		temp = temp.next
	}

	if temp != nil {
		temp.value = value
		return
	}

	prev.next = &node{
		value: value,
		key:   key,
	}
}

func (m *MyHashMap) Get(key int) int {
	calcKey := m.hash(key)

	if m.data[calcKey] == nil {
		return -1
	}

	temp := m.data[calcKey]
	for temp != nil {
		if temp.key == key {
			return temp.value
		}
		temp = temp.next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	calcKey := m.hash(key)

	if m.data[calcKey] == nil {
		return
	}

	var prev *node
	temp := m.data[calcKey]
	for temp != nil && temp.key != key {
		prev = temp
		temp = temp.next
	}

	if temp != nil {
		if prev == nil {
			m.data[calcKey] = temp.next
		} else {
			prev.next = temp.next
		}
	}
}

func (m *MyHashMap) hash(key int) int {
	return key % m.size
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
