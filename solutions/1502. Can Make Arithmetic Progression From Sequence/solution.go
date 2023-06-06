func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr)

	min := arr[0]
	max := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if (max-min)%(n-1) != 0 {
		return false
	}

	d := (max - min) / (n - 1)
	if d == 0 {
		return true
	}

	if (max-min)%d != 0 {
		return false
	}

	counts := make([]bool, (max-min)/d+1)
	for _, v := range arr {
		if (v-min)%d != 0 {
			return false
		}

		seqNumber := (v - min) / d
		if counts[seqNumber] {
			return false
		}

		counts[seqNumber] = true
	}

	for _, v := range counts {
		if !v {
			return false
		}
	}

	return true
}
