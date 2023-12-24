import "math"

func minOperations(s string) int {
	count1 := 0 // 010
	count2 := 0 // 101

	for i := range s {
		if i%2 == 0 {
			if s[i] == '1' {
				count1++
			} else {
				count2++
			}
		} else {
			if s[i] == '1' {
				count2++
			} else {
				count1++
			}
		}
	}

	return int(math.Min(float64(count1), float64(count2)))
}
