func canArrange(arr []int, k int) bool {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[makeKey(v, k)]++
	}

	for _, v := range arr {
		need := makeKey(v, k)

		if need == 0 {
			if counts[need]%2 == 1 {
				return false
			}
		} else if counts[need] != counts[k-need] {
			return false
		}
	}

	return true
}

func makeKey(v, k int) int {
	return (v%k + k) % k
}
