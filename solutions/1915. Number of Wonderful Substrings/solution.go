func wonderfulSubstrings(word string) int64 {
	cnt := make([]int64, 1024)
	cnt[0] = 1

	curState := 0
	var res int64

	for _, c := range word {
		curState ^= 1 << get(c)

		res += cnt[curState]
		for odd := 'a'; odd <= 'j'; odd++ {
			oddState := curState ^ (1 << get(odd))
			res += cnt[oddState]
		}

		cnt[curState]++
	}

	return res
}

func get(c rune) int {
	return int(c - 'a')
}
