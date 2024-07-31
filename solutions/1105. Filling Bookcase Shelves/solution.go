import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	cache := make(map[int]int)

	var helper func(i int) int
	helper = func(i int) int {
		if i == len(books) {
			return 0
		}
		if val, ok := cache[i]; ok {
			return val
		}

		curWidth := shelfWidth
		maxHeight := 0
		cache[i] = math.MaxInt

		for j := i; j < len(books); j++ {
			width, height := books[j][0], books[j][1]
			if width > curWidth {
				break
			}
			curWidth -= width
			maxHeight = max(maxHeight, height)
			cache[i] = min(cache[i], helper(j+1)+maxHeight)
		}

		return cache[i]
	}

	return helper(0)
}
