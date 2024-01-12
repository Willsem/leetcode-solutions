func isVowel(c rune) int {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return 1
	default:
		return 0
	}
}

func halvesAreAlike(s string) bool {
	count1, count2 := 0, 0
	for i, v := range s {
		if i < len(s)/2 {
			count1 += isVowel(v)
		} else {
			count2 += isVowel(v)
		}
	}

	return count1 == count2
}
