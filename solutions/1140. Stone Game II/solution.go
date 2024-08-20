func stoneGameII(piles []int) int {
	totalPiles := len(piles)
	suffixSums := make([]int, totalPiles+1)
	for i := totalPiles - 1; i >= 0; i-- {
		suffixSums[i] = suffixSums[i+1] + piles[i]
	}

	memo := make([][]int, totalPiles)
	for i := range memo {
		memo[i] = make([]int, totalPiles+1)
	}

	var maxStonesAliceCanGet func(int, int) int
	maxStonesAliceCanGet = func(m, currentPile int) int {
		if currentPile >= totalPiles {
			return 0
		}

		if currentPile+2*m >= totalPiles {
			return suffixSums[currentPile]
		}

		if memo[currentPile][m] != 0 {
			return memo[currentPile][m]
		}

		maxStones := 0

		for x := 1; x <= 2*m; x++ {
			currentStones := suffixSums[currentPile] - maxStonesAliceCanGet(max(m, x), currentPile+x)
			maxStones = max(maxStones, currentStones)
		}

		memo[currentPile][m] = maxStones
		return maxStones
	}

	return maxStonesAliceCanGet(1, 0)
}
