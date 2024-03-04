import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	max, currPoints := 0, 0
	left, right := 0, len(tokens)-1
	for left <= right {
		if power >= tokens[left] {
			currPoints++
			if currPoints > max {
				max = currPoints
			}
			power -= tokens[left]
			left++
			continue
		}

		if currPoints == 0 {
			break
		}

		currPoints--
		power += tokens[right]
		right--
	}

	return max
}
