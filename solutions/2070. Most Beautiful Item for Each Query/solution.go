import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	ans := make([]int, 0, len(queries))

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	maxBeauty := items[0][1]
	for i := 0; i < len(items); i++ {
		maxBeauty = max(maxBeauty, items[i][1])
		items[i][1] = maxBeauty
	}

	for i := 0; i < len(queries); i++ {
		ans = append(ans, binarySearch(items, queries[i]))
	}

	return ans
}

func binarySearch(items [][]int, targetPrice int) int {
	l, r := 0, len(items)-1
	maxBeauty := 0

	for l <= r {
		mid := (l + r) / 2
		if items[mid][0] > targetPrice {
			r = mid - 1
		} else {
			maxBeauty = max(maxBeauty, items[mid][1])
			l = mid + 1
		}
	}

	return maxBeauty
}
