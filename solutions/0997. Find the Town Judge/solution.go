func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	countTrustIn := make(map[int]int)
	countTrustOut := make(map[int]int)
	for _, t := range trust {
		countTrustIn[t[1]]++
		countTrustOut[t[0]]++

		if _, ok := countTrustIn[t[0]]; !ok {
			countTrustIn[t[0]] = 0
		}
		if _, ok := countTrustOut[t[1]]; !ok {
			countTrustOut[t[1]] = 0
		}
	}

	for citizen := range countTrustIn {
		if countTrustIn[citizen] == n-1 && countTrustOut[citizen] == 0 {
			return citizen
		}
	}

	return -1
}
