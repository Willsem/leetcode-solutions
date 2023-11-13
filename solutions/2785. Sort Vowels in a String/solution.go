var vowels = []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

func sortVowels(s string) string {
	vowelsCounts := make(map[byte]int, len(vowels))
	for _, vowel := range vowels {
		vowelsCounts[vowel] = 0
	}

	newS := []byte(s)
	for i, v := range newS {
		for _, vowel := range vowels {
			if v == vowel {
				vowelsCounts[vowel]++
				newS[i] = 'a'
			}
		}
	}

	curr := 0
	for i := range newS {
		if newS[i] == 'a' {
			for curr < len(vowels) && vowelsCounts[vowels[curr]] == 0 {
				curr++
			}

			if curr == len(vowels) {
				return string(newS)
			}

			newS[i] = vowels[curr]
			vowelsCounts[vowels[curr]]--
		}
	}

	return string(newS)
}
