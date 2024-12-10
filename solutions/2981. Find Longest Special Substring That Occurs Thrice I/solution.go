func maximumLength(s string) int {
	mp := [26][]int{}
	ans := -1
	preLetter := ' '

	for _, a := range s {
		a -= 97
		if preLetter == a {
			mp[a][len(mp[a])-1]++
		} else {
			mp[a] = append(mp[a], 1)
		}
		preLetter = a
	}

	for _, a := range mp {
		subCountMap := make(map[int]int)
		for _, count := range a {
			for i := 1; count >= 1; count-- {
				subCountMap[count] += i
				i++
			}
		}
		for maxSubLen, count := range subCountMap {
			if count >= 3 {
				ans = max(ans, maxSubLen)
			}
		}
	}

	return ans
}
