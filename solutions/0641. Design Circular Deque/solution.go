type MyCircularDeque struct {
	queue    []int
	front    int
	rear     int
	size     int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue:    make([]int, k),
		front:    0,
		rear:     k - 1,
		size:     0,
		capacity: k,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}

	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.queue[d.front] = value
	d.size++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}

	d.rear = (d.rear + 1) % d.capacity
	d.queue[d.rear] = value
	d.size++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}

	d.front = (d.front + 1) % d.capacity
	d.size--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}

	d.rear = (d.rear - 1 + d.capacity) % d.capacity
	d.size--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}

	return d.queue[d.front]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}

	return d.queue[d.rear]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
