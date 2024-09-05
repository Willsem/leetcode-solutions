func missingRolls(rolls []int, mean int, n int) []int {
	currSum := 0
	for _, roll := range rolls {
		currSum += roll
	}

	needSum := mean * (len(rolls) + n)
	requiredSum := needSum - currSum

	if requiredSum > 6*n || requiredSum < n {
		return []int{}
	}

	requiredMean := requiredSum / n
	mod := requiredSum % n

	result := make([]int, n)
	for i := range result {
		result[i] = requiredMean
	}
	for i := 0; i < mod; i++ {
		result[i]++
	}

	return result
}
