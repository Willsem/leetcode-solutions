func getRow(rowIndex int) []int {
	res := []int{1}

	prev := 1
	for i := 1; i <= rowIndex; i++ {
		next := prev * (rowIndex - i + 1) / i
		res = append(res, next)
		prev = next
	}

	return res
}
