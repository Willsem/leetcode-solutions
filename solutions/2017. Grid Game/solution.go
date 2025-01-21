import "math"

func gridGame(grid [][]int) int64 {
	firstRowSum := 0
	for _, num := range grid[0] {
		firstRowSum += num
	}

	secondRowSum := 0
	minSum := math.MaxInt
	for i := 0; i < len(grid[0]); i++ {
		firstRowSum -= grid[0][i]
		minSum = min(minSum, max(firstRowSum, secondRowSum))
		secondRowSum += grid[1][i]
	}

	return int64(minSum)
}
