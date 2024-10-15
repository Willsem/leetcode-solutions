func minimumSteps(s string) int64 {
	swaps, blacks := 0, 0
	for _, c := range s {
		if c == '0' {
			swaps += blacks
		} else {
			blacks++
		}
	}
	return int64(swaps)
}
