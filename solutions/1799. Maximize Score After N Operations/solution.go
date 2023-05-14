func maxScore(nums []int) int {
	n := len(nums)

	gcdMat := make([][]int, n)
	for i := 1; i < n; i++ {
		gcdMat[i] = make([]int, i)
		for j := 0; j < i; j++ {
			gcdMat[i][j] = gcd(nums[i], nums[j])
		}
	}

	statesAmount := 1 << n
	allPickedState := statesAmount - 1

	dp := make([]int, statesAmount)
	dp[0] = 0

	for state := 1; state <= allPickedState; state++ {
		pickedNumbersAmount := getPickedNumbersAmount(state)
		pickedPairsAmount := pickedNumbersAmount / 2

		if pickedNumbersAmount%2 == 1 {
			continue
		}

		for i := 1; i < n; i++ {
			if state&(1<<i) == 0 {
				continue
			}

			for j := 0; j < i; j++ {
				if state&(1<<j) == 0 {
					continue
				}

				stateWithoutIJ := state & ^(1 << i) & ^(1 << j)
				score := dp[stateWithoutIJ] + pickedPairsAmount*gcdMat[i][j]
				dp[state] = max(dp[state], score)
			}
		}
	}
	return dp[allPickedState]
}

func getPickedNumbersAmount(state int) int {
	amount := 0
	for state > 0 {
		amount++
		state &= state - 1
	}
	return amount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
