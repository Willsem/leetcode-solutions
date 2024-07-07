func numWaterBottles(numBottles int, numExchange int) int {
	ans := 0
	ost := 0
	for numBottles+ost >= numExchange {
		ans += numBottles

		if (numBottles+ost)/numExchange > numBottles/numExchange {
			numBottles += ost
			ost = 0
		}

		ost += numBottles % numExchange
		numBottles /= numExchange
	}

	return ans + numBottles
}
