func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	res := 0
	time := 0
	for i := range startTime {
		time += endTime[i] - startTime[i]

		left := 0
		if i > k-1 {
			left = endTime[i-k]
		}

		right := eventTime
		if i < len(startTime)-1 {
			right = startTime[i+1]
		}

		if right-left-time > res {
			res = right - left - time
		}

		if i >= k-1 {
			time -= endTime[i-k+1] - startTime[i-k+1]
		}
	}

	return res
}
