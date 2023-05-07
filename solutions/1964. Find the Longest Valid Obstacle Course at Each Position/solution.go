func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	dp := make([]int, 0, len(obstacles))
	answer := make([]int, len(obstacles))

	for i := range obstacles {
		length := search(dp, obstacles[i]+1)
		if length == len(dp) {
			dp = append(dp, obstacles[i])
		} else {
			dp[length] = obstacles[i]
		}

		answer[i] = length + 1
	}

	return answer
}

func search(arr []int, target int) int {
	l := -1
	r := len(arr)

	for l < r-1 {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}
