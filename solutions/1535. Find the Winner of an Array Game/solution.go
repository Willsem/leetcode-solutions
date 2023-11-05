func getWinner(arr []int, k int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	curr := arr[0]
	winStreak := 0

	for i := 1; i < len(arr); i++ {
		op := arr[i]

		if curr > op {
			winStreak++
		} else {
			curr = op
			winStreak = 1
		}

		if winStreak == k || curr == max {
			return curr
		}
	}

	return -1
}
