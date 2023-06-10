func maxValue(n int, index int, maxSum int) int {
	if n == maxSum {
		return 1
	}

	ans := (maxSum + (index*(index+1)+(n-1-index)*(n-index))/2) / n

	if ans > index && ans > n-1-index {
		// maxSum is large enough
		return ans
	}

	total := n
	left, right := 0, 0
	ans = 1
	for total < maxSum {
		ans++
		if left < index {
			left++
		}
		if right < n-1-index {
			right++
		}
		total += left + right + 1
	}

	return ans
}
