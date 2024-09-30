type CustomStack struct {
	stack      []int
	increments []int
	capacity   int
	curr       int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack:      make([]int, maxSize),
		increments: make([]int, maxSize),
		capacity:   maxSize,
		curr:       -1,
	}
}

func (this *CustomStack) Push(x int) {
	if this.curr == this.capacity-1 {
		return
	}

	this.curr++
	this.stack[this.curr] = x
}

func (this *CustomStack) Pop() int {
	if this.curr < 0 {
		return -1
	}

	element := this.stack[this.curr] + this.increments[this.curr]
	this.curr--
	if this.curr > -1 {
		this.increments[this.curr] += this.increments[this.curr+1]
	}
	this.increments[this.curr+1] = 0

	return element
}

func (this *CustomStack) Increment(k int, val int) {
	if this.curr < 0 {
		return
	}

	k--
	if k > this.curr {
		k = this.curr
	}

	this.increments[k] += val
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
