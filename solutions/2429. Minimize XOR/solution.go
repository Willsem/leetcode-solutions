import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	bitsCount := bits.OnesCount(uint(num2))

	currBitsCount := 0
	currBit := 31

	result := 0
	for currBitsCount < bitsCount {
		if isSet(num1, currBit) || bitsCount-currBitsCount > currBit {
			result = setBit(result, currBit)
			currBitsCount++
		}
		currBit--
	}

	return result
}

func isSet(num, bit int) bool {
	return num&(1<<bit) > 0
}

func setBit(num, bit int) int {
	return num | (1 << bit)
}
