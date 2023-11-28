const mod = 1e9 + 7

func numberOfWays(corridor string) int {
	count := 1
	seats := 0
	prevPairLast := -1

	for i := 0; i < len(corridor); i++ {
		if corridor[i] == 'S' {
			seats++
			if seats == 2 {
				prevPairLast = i
				seats = 0
			} else if seats == 1 && prevPairLast != -1 {
				count *= (i - prevPairLast)
				count %= mod
			}
		}
	}

	if seats == 1 || prevPairLast == -1 {
		return 0
	}

	return count
}
