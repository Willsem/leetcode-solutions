import (
	"sort"
	"strconv"
)

const digits = "123456789"

func sequentialDigits(low int, high int) []int {
	ans := make([]int, 0)

	for i := 0; i < len(digits); i++ {
		for j := i + 1; j <= len(digits); j++ {
			curr, _ := strconv.Atoi(digits[i:j])
			if curr >= low && curr <= high {
				ans = append(ans, curr)
			}
		}
	}

	sort.Ints(ans)
	return ans
}
