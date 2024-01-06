import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	dp := [][]int{{0, 0}}

	for _, job := range jobs {
		start, end, prof := job[0], job[1], job[2]
		i := binarySearch(0, len(dp), func(i int) bool {
			return dp[i][0] <= start
		}, false)
		newProfit := dp[i-1][1] + prof
		if newProfit > dp[len(dp)-1][1] {
			dp = append(dp, []int{end, newProfit})
		}
	}
	return dp[len(dp)-1][1]
}

func binarySearch(l, r int, comp func(i int) bool, leftSearch bool) int {
	for l < r {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if leftSearch {
		return r
	}

	return l
}
