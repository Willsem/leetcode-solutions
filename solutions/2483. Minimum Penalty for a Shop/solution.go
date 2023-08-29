func bestClosingTime(customers string) int {
	penalty := 0
	for _, c := range customers {
		if c == 'Y' {
			penalty++
		}
	}

	minPenalty := penalty
	minHour := 0

	for i, c := range customers {
		if c == 'Y' {
			penalty--
		} else {
			penalty++
		}

		if penalty < minPenalty {
			minPenalty = penalty
			minHour = i + 1
		}
	}

	return minHour
}
