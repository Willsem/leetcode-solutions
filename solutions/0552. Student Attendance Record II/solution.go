const mod = 1000000007

func checkRecord(n int) int {
	dpCurrState := make([][]int, 2)
	for i := range dpCurrState {
		dpCurrState[i] = make([]int, 3)
	}

	dpNextState := make([][]int, 2)
	for i := range dpNextState {
		dpNextState[i] = make([]int, 3)
	}

	dpCurrState[0][0] = 1
	for len := 0; len < n; len++ {
		for totalAbsences := 0; totalAbsences <= 1; totalAbsences++ {
			for consecutiveLates := 0; consecutiveLates <= 2; consecutiveLates++ {
				dpNextState[totalAbsences][0] = (dpNextState[totalAbsences][0] + dpCurrState[totalAbsences][consecutiveLates]) % mod
				if totalAbsences < 1 {
					dpNextState[totalAbsences+1][0] = (dpNextState[totalAbsences+1][0] + dpCurrState[totalAbsences][consecutiveLates]) % mod
				}
				if consecutiveLates < 2 {
					dpNextState[totalAbsences][consecutiveLates+1] = (dpNextState[totalAbsences][consecutiveLates+1] + dpCurrState[totalAbsences][consecutiveLates]) % mod
				}
			}
		}

		dpCurrState, dpNextState = dpNextState, dpCurrState
		for i := range dpNextState {
			for j := range dpNextState[i] {
				dpNextState[i][j] = 0
			}
		}
	}

	count := 0
	for totalAbsences := 0; totalAbsences <= 1; totalAbsences++ {
		for consecutiveLates := 0; consecutiveLates <= 2; consecutiveLates++ {
			count = (count + dpCurrState[totalAbsences][consecutiveLates]) % mod
		}
	}

	return count
}
