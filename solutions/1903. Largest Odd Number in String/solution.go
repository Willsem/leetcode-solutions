func largestOddNumber(num string) string {
	found := -1
	for i := len(num) - 1; i >= 0; i-- {
		if num[i]%2 == 1 {
			found = i
			break
		}
	}

	if found == -1 {
		return ""
	}

	return num[:found+1]
}
