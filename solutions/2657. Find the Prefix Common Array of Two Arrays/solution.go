func findThePrefixCommonArray(A []int, B []int) []int {
	numbersA := make(map[int]struct{})
	numbersB := make(map[int]struct{})
	intersection := make(map[int]struct{})

	result := make([]int, len(A))
	for i := range A {
		numbersA[A[i]] = struct{}{}
		numbersB[B[i]] = struct{}{}

		if _, ok := numbersB[A[i]]; ok {
			intersection[A[i]] = struct{}{}
		}

		if _, ok := numbersA[B[i]]; ok {
			intersection[B[i]] = struct{}{}
		}

		result[i] = len(intersection)
	}

	return result
}
