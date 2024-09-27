type MyCalendarTwo struct {
	single [][]int
	double [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		single: make([][]int, 0),
		double: make([][]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, event := range this.double {
		if max(start, event[0]) < min(end, event[1]) {
			return false
		}
	}

	for _, event := range this.single {
		if max(start, event[0]) < min(end, event[1]) {
			this.double = append(this.double, []int{max(start, event[0]), min(end, event[1])})
		}
	}

	this.single = append(this.single, []int{start, end})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
