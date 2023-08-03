var mapping = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)

	var solve func(path string)
	solve = func(path string) {
		if len(path) == len(digits) {
			result = append(result, path)
			return
		}

		curr := len(path)
		digit := digits[curr]
		for _, sym := range mapping[digit] {
			solve(path + string(sym))
		}
	}

	if len(digits) > 0 {
		solve("")
	}
	return result
}
