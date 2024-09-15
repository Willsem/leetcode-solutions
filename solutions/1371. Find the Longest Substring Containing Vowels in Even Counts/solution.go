func findTheLongestSubstring(s string) int {
	charMap := make([]int, 26)
	charMap['a'-'a'] = 1
	charMap['e'-'a'] = 2
	charMap['i'-'a'] = 4
	charMap['o'-'a'] = 8
	charMap['u'-'a'] = 16

	prevXOR := make([]int, 32)
	for i := range prevXOR {
		prevXOR[i] = -1
	}

	prefixXOR := 0
	result := 0
	for i := range s {
		prefixXOR ^= charMap[s[i]-'a']
		if prevXOR[prefixXOR] == -1 && prefixXOR != 0 {
			prevXOR[prefixXOR] = i
		}
		result = max(result, i-prevXOR[prefixXOR])
	}

	return result
}
