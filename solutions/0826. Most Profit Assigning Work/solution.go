func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	maxWorker := worker[0]
	for _, w := range worker {
		if maxWorker < w {
			maxWorker = w
		}
	}

	profits := make([]int, maxWorker+1)
	for i, diff := range difficulty {
		if diff <= maxWorker {
			profits[diff] = max(profits[diff], profit[i])
		}
	}

	for i := 1; i < len(profits); i++ {
		profits[i] = max(profits[i], profits[i-1])
	}

	res := 0
	for _, w := range worker {
		res += profits[w]
	}
	return res
}
