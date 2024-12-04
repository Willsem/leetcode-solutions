func canMakeSubsequence(str1 string, str2 string) bool {
	i2 := 0

	for i1 := 0; i1 < len(str1) && i2 < len(str2); i1++ {
		if str1[i1] == str2[i2] || str1[i1]+1 == str2[i2] || str1[i1]-25 == str2[i2] {
			i2++
		}
	}

	return i2 == len(str2)
}
