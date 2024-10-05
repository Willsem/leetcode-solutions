func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	left, right := 0, len(s1)-1

	count1 := make(map[byte]int)
	count2 := make(map[byte]int)
	for i := left; i <= right; i++ {
		count1[s1[i]]++
		count2[s2[i]]++
	}

	for right < len(s2) {
		equals := true
		if len(count1) != len(count2) {
			equals = false
		} else {
			for k, v := range count1 {
				if count2[k] != v {
					equals = false
					break
				}
			}
		}

		if equals {
			return true
		}

		count2[s2[left]]--
		if count2[s2[left]] == 0 {
			delete(count2, s2[left])
		}

		left++
		right++

		if right < len(s2) {
			count2[s2[right]]++
		}
	}

	return false
}
