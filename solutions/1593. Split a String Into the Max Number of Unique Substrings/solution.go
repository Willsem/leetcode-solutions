func maxUniqueSplit(s string) int {
	was := make(map[string]struct{})
	return maxUniqueSplitRecursive(s, 0, was)
}

func maxUniqueSplitRecursive(s string, start int, was map[string]struct{}) int {
	if start == len(s) {
		return 0
	}

	maxCount := 0
	for end := start + 1; end <= len(s); end++ {
		substr := s[start:end]
		if _, ok := was[substr]; !ok {
			was[substr] = struct{}{}
			maxCount = max(maxCount, maxUniqueSplitRecursive(s, end, was)+1)
			delete(was, substr)
		}
	}

	return maxCount
}
