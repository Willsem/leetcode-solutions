func minDominoRotations(tops []int, bottoms []int) int {
	countTop := make([]int, 6)
	countBot := make([]int, 6)
	same := make([]int, 6)

	for i := range tops {
		countTop[tops[i]-1]++
		countBot[bottoms[i]-1]++
		if tops[i] == bottoms[i] {
			same[tops[i]-1]++
		}
	}

	for i := range countTop {
		if countTop[i]+countBot[i]-same[i] == len(tops) {
			if countTop[i] < countBot[i] {
				return countTop[i] - same[i]
			}
			return countBot[i] - same[i]
		}
	}

	return -1
}
