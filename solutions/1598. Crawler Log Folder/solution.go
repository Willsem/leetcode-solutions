func minOperations(logs []string) int {
	deep := 0

	for _, log := range logs {
		switch log {
		case "./":
		case "../":
			if deep > 0 {
				deep--
			}
		default:
			deep++
		}
	}

	return deep
}
