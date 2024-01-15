import "sort"

func findWinners(matches [][]int) [][]int {
	countLoses := make(map[int]int)
	for _, match := range matches {
		winner, loser := match[0], match[1]
		countLoses[loser]++
		if _, ok := countLoses[winner]; !ok {
			countLoses[winner] = 0
		}
	}

	ans := [][]int{
		make([]int, 0),
		make([]int, 0),
	}
	for player, count := range countLoses {
		switch count {
		case 0:
			ans[0] = append(ans[0], player)
		case 1:
			ans[1] = append(ans[1], player)
		}
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}
