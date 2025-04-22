func idealArrays(n int, maxValue int) int {
	mod := int64(1_000_000_007)
	maxLen := computeMaxLen(maxValue)
	multiplerTable := computeMutiplierTable(n-1, maxLen, mod)

	dp := make(map[int]int64)

	for lastValue := 1; lastValue <= maxValue; lastValue++ {
		dp[lastValue] = 1
	}
	result := int64(maxValue)

	for len := 2; len <= maxLen; len++ {
		prevDp := dp
		dp = make(map[int]int64)
		for lastValue, count := range prevDp {
			for value := lastValue * 2; value <= maxValue; value += lastValue {
				dp[value] = (dp[value] + count) % mod
			}
		}

		subTotal := int64(0)
		for _, count := range dp {
			subTotal = (subTotal + count) % mod
		}
		result = (result + multiplerTable[len-1]*subTotal) % mod
	}
	return int(result)
}

func computeMaxLen(maxValue int) int {
	result := 0
	num := maxValue
	for num > 0 {
		num /= 2
		result++
	}
	return result
}

func computeMutiplierTable(n, maxLen int, mod int64) []int64 {
	pendingFactors := make(map[int]int)
	for k := 2; k < maxLen; k++ {
		for _, factorPowerPair := range factorize(k) {
			pendingFactors[factorPowerPair[0]] += factorPowerPair[1]
		}
	}

	result := make([]int64, maxLen)
	result[0] = 1
	multiplierL := int64(1)
	multiplierR := int64(1)

	for k := 1; k < maxLen; k++ {
		numL := 1
		numR := n - k + 1
		for factor, count := range pendingFactors {
			if count == 0 {
				continue
			}
			remainingCount := count
			for remainingCount > 0 && numR > 1 && numR%factor == 0 {
				remainingCount--
				numL *= factor
				numR /= factor
			}
			pendingFactors[factor] = remainingCount
		}

		multiplierL = (multiplierL * int64(numL)) / int64(k)
		multiplierR = (multiplierR * int64(numR)) % mod
		result[k] = (multiplierL * multiplierR) % mod
	}
	return result
}

func factorize(num int) [][]int {
	result := make([][]int, 0)

	x := num
	for factor := 2; factor <= x; factor++ {
		pow := 0
		for x%factor == 0 {
			pow++
			x /= factor
		}
		if pow > 0 {
			result = append(result, []int{factor, pow})
		}
	}
	return result
}
