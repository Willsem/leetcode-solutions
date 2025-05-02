func pushDominoes(dominoes string) string {
	if len(dominoes) == 1 {
		return dominoes
	}

	result := []byte(dominoes)
	for i := range result {
		if result[i] == '.' {
			left := false
			if i > 0 {
				if result[i-1] == 'R' {
					left = true
				}
			}

			right, count := false, 0
			for j := i; j < len(result) && result[j] == '.'; j++ {
				count++
			}

			if i+count < len(result) {
				if result[i+count] == 'L' {
					right = true
				}
			}

			if left && right {
				for j := i; j < i+count/2; j++ {
					result[j] = 'R'
				}
				for j := i + (count+1)/2; j < len(result) && j < i+count; j++ {
					result[j] = 'L'
				}
			} else if right {
				for j := i; j < len(result) && j < i+count; j++ {
					result[j] = 'L'
				}
			} else if left {
				for j := i; j < len(result) && j < i+count; j++ {
					result[j] = 'R'
				}
			}

			i = i + count - 1
		}
	}

	return string(result)
}
