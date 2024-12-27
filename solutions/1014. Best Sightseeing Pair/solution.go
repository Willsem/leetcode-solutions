func maxScoreSightseeingPair(values []int) int {
	maxScore := 0
	bestLeft := values[0]

	for i := 1; i < len(values); i++ {
		maxScore = max(maxScore, bestLeft+values[i]-i)
		bestLeft = max(bestLeft, values[i]+i)
	}

	return maxScore
}
