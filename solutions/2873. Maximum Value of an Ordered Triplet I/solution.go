func maximumTripletValue(nums []int) int64 {
	result, imax, dmax := 0, 0, 0
	for _, num := range nums {
		result = max(result, dmax*num)
		dmax = max(dmax, imax-num)
		imax = max(imax, num)
	}
	return int64(result)
}
