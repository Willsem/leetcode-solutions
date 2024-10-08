func minSwaps(s string) int {
	stackSize := 0

	for _, c := range s {
		if c == '[' {
			stackSize++
		} else if stackSize > 0 {
			stackSize--
		}
	}

	return (stackSize + 1) / 2
}
