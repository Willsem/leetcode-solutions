import "sort"

func closeStrings(word1 string, word2 string) bool {
	counts1, counts2 := make([]int, 26), make([]int, 26)

	for _, v := range word1 {
		counts1[int(v)-int('a')]++
	}
	for _, v := range word2 {
		counts2[int(v)-int('a')]++
	}

	for i := 0; i < 26; i++ {
		if (counts1[i] == 0 && counts2[i] != 0) ||
			(counts1[i] != 0 && counts2[i] == 0) {
			return false
		}
	}

	sort.Ints(counts1)
	sort.Ints(counts2)

	for i := 0; i < 26; i++ {
		if counts1[i] != counts2[i] {
			return false
		}
	}

	return true
}
