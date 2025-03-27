func minimumIndex(nums []int) int {
	freq := make(map[int]int)
	x := nums[0]
	for _, num := range nums {
		freq[num]++

		if freq[num] > len(nums)/2 {
			x = num
		}
	}

	xFreq := freq[x]
	leftArrFreq, rightArrFreq := 0, xFreq
	for i := range nums {
		if nums[i] == x {
			leftArrFreq++
			rightArrFreq--
		}

		if leftArrFreq*2 > i+1 && rightArrFreq*2 > len(nums)-i-1 {
			return i
		}
	}

	return -1
}
