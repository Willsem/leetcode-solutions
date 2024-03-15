func productExceptSelf(nums []int) []int {
	n := len(nums)

	preProd := make([]int, n)
	preProd[0] = 1
	for i := 1; i < n; i++ {
		preProd[i] = preProd[i-1] * nums[i-1]
	}

	postProd := make([]int, n)
	postProd[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		postProd[i] = postProd[i+1] * nums[i+1]
	}

	for i := range preProd {
		preProd[i] *= postProd[i]
	}

	return preProd
}
