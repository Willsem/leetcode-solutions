func differenceOfSums(n int, m int) int {
	k := n / m
	num2 := k * (k + 1) * m / 2
	num1 := n*(n+1)/2 - num2
	return num1 - num2
}
