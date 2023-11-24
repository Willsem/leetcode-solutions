import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)

	ans := 0
	for i := len(piles) / 3; i < len(piles); i += 2 {
		ans += piles[i]
	}

	return ans
}
