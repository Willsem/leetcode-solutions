import "sort"

func findEvenNumbers(digits []int) []int {
	mapResult := make(map[int]struct{})

	for i, di := range digits {
		if di == 0 {
			continue
		}

		for j, dj := range digits {
			if i == j {
				continue
			}

			for k, dk := range digits {
				if i == k || j == k || digits[k]%2 != 0 {
					continue
				}

				mapResult[di*100+dj*10+dk] = struct{}{}
			}
		}
	}

	result := make([]int, 0, len(mapResult))
	for num := range mapResult {
		result = append(result, num)
	}
	sort.Ints(result)

	return result
}
