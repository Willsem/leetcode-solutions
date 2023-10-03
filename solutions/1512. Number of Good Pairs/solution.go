func numIdenticalPairs(nums []int) int {
	prev := make(map[int]int)
	count := 0

	for _, v := range nums {
		if _, ok := prev[v]; !ok {
			prev[v] = 1
		} else {
			count += prev[v]
			prev[v]++
		}
	}

	return count
}
