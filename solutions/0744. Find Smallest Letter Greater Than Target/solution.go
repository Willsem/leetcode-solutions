func nextGreatestLetter(letters []byte, target byte) byte {
	index := binarySearch(0, len(letters), func(i int) bool {
		return letters[i] <= target
	})

	if index+1 >= len(letters) {
		return letters[0]
	}

	if letters[index] > target {
		return letters[index]
	}

	return letters[index+1]
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r-1 {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}
