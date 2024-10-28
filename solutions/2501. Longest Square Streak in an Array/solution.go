func longestSquareStreak(nums []int) int {
	maxStreak := 0

	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	for _, num := range nums {
		currStreak := 0

		for _, ok := set[num]; ok && num <= 1e5; _, ok = set[num] {
			currStreak++
			num *= num
		}

		maxStreak = max(maxStreak, currStreak)
	}

	if maxStreak < 2 {
		return -1
	}

	return maxStreak
}
