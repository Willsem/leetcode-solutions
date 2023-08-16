func maxSlidingWindow(nums []int, k int) []int {
	var res []int

	var queue MonotonicQueue

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			queue.push(nums[i])
		} else {
			queue.push(nums[i])
			res = append(res, queue[0])
			queue.pop(nums[i-k+1])
		}
	}

	return res
}

type MonotonicQueue []int

func (m *MonotonicQueue) length() int {
	return len(*m)
}

func (m *MonotonicQueue) back() int {
	if m.length() == 0 {
		panic("empty queue")
	}

	return (*m)[len(*m)-1]
}
func (m *MonotonicQueue) pop(item int) {
	if m.length() > 0 && (*m)[0] == item {
		*m = (*m)[1:]
	}
}

func (m *MonotonicQueue) popBack() int {
	if len(*m) == 0 {
		panic("empty queue")
	}

	poppedItem := (*m)[len(*m)-1]

	*m = (*m)[:len(*m)-1]

	return poppedItem
}

func (m *MonotonicQueue) push(item int) {
	for m.length() > 0 && m.back() < item {
		m.popBack()
	}

	*m = append(*m, item)
}
