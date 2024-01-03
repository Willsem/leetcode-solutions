func numberOfBeams(bank []string) int {
	res, prev := 0, 0

	for _, row := range bank {
		count := 0
		for _, c := range row {
			if c == '1' {
				count++
			}
		}
		if count != 0 {
			res += prev * count
			prev = count
		}
	}

	return res
}
