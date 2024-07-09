func averageWaitingTime(customers [][]int) float64 {
	waitTime, next := 0, 0
	for _, c := range customers {
		next = max(next, c[0]) + c[1]
		waitTime += next - c[0]
	}
	return float64(waitTime) / float64(len(customers))
}
