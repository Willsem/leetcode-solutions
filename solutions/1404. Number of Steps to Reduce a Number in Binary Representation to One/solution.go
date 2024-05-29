func numSteps(s string) int {
	count, carry := 0, 0
	for i := len(s) - 1; i > 0; i-- {
		curr := int(s[i]) - int('0')
		if (curr+carry)%2 == 1 {
			count += 2
			carry = 1
		} else {
			count++
		}
	}

	return count + carry
}
