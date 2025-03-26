import "sort"

func minOperations(grid [][]int, x int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 0
	}

	remainder := grid[0][0] % x
	for _, row := range grid {
		for _, el := range row {
			if el%x != remainder {
				return -1
			}
		}
	}

	arr := make([]int, 0)
	for _, row := range grid {
		for _, el := range row {
			arr = append(arr, el)
		}
	}

	sort.Ints(arr)

	median1, median2 := arr[len(arr)/2], arr[len(arr)/2-1]
	if len(arr)%2 == 1 {
		median2 = median1
	}

	count1, count2 := 0, 0
	for _, el := range arr {
		count1 += abs(median1-el) / x
		count2 += abs(median2-el) / x
	}

	return min(count1, count2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
