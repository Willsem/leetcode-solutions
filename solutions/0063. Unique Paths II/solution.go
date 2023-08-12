func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := range obstacleGrid {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
			} else if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			}
		}
	}

	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
