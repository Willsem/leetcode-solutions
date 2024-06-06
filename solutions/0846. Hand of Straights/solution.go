import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	counts := make(map[int]int)
	for _, card := range hand {
		counts[card]++
	}

	splits := 0
	for _, card := range hand {
		if counts[card] > 0 {
			splits++
			for i := 0; i < groupSize; i++ {
				if counts[card+i] == 0 {
					return false
				} else {
					counts[card+i]--
				}
			}
		}
	}

	return splits == len(hand)/groupSize
}
