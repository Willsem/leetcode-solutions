func findKthBit(n int, k int) byte {
	posInSection := k & (-k)
	isInInvertedPart := ((k / posInSection) >> 1 & 1) == 1
	originalBitIsOne := (k & 1) == 0
	if isInInvertedPart {
		if originalBitIsOne {
			return '0'
		} else {
			return '1'
		}
	} else {
		if originalBitIsOne {
			return '1'
		} else {
			return '0'
		}
	}
}
