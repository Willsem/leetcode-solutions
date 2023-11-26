import "math"

type data struct {
	height, col int
}

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	prevHeights := make([]data, 0)
	ans := 0

	for row := 0; row < m; row++ {
		heights := make([]data, 0)
		seen := make([]bool, n)

		for _, height := range prevHeights {
			height, col := height.height, height.col
			if matrix[row][col] == 1 {
				heights = append(heights, data{height + 1, col})
				seen[col] = true
			}
		}

		for col := 0; col < n; col++ {
			if !seen[col] && matrix[row][col] == 1 {
				heights = append(heights, data{1, col})
			}
		}

		for i := 0; i < len(heights); i++ {
			ans = int(math.Max(float64(ans), float64(heights[i].height*(i+1))))
		}

		prevHeights = heights
	}

	return ans
}
