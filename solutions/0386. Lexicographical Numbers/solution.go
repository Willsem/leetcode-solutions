func lexicalOrder(n int) []int {
	result := make([]int, 0, n)
	for i, curr := 0, 1; i < n; i++ {
		result = append(result, curr)
		if curr*10 <= n {
			curr *= 10
		} else {
			for curr%10 == 9 || curr >= n {
				curr /= 10
			}
			curr++
		}
	}
	return result
}
