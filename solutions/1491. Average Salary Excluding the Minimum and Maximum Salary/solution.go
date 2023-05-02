func average(salary []int) float64 {
	max := salary[0]
	min := salary[0]
	sum := 0

	for _, v := range salary {
		sum += v
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return float64(sum-max-min) / float64(len(salary)-2)
}
