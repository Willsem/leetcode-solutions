const mod = 1e9 + 7

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := makeMatrix(m, n)

	dp[startRow][startColumn] = 1
	count := 0

	for moves := 1; moves <= maxMove; moves++ {
		temp := makeMatrix(m, n)

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == m-1 {
					count = (count + dp[i][j]) % mod
				}
				if j == n-1 {
					count = (count + dp[i][j]) % mod
				}
				if i == 0 {
					count = (count + dp[i][j]) % mod
				}
				if j == 0 {
					count = (count + dp[i][j]) % mod
				}

				t1 := 0
				if i > 0 {
					t1 = dp[i-1][j]
				}
				t2 := 0
				if i < m-1 {
					t2 = dp[i+1][j]
				}
				t3 := 0
				if j > 0 {
					t3 = dp[i][j-1]
				}
				t4 := 0
				if j < n-1 {
					t4 = dp[i][j+1]
				}

				temp[i][j] = ((t1+t2)%mod + (t3+t4)%mod) % mod
			}
		}

		dp = temp
	}

	return count
}

func makeMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}
