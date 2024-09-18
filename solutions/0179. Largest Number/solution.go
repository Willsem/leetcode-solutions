import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := fmt.Sprint(nums[i]), fmt.Sprint(nums[j])
		return !comp(a, b)
	})

	if nums[0] == 0 {
		return "0"
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), ""), "[]")
}

func comp(a, b string) bool {
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}

	if len(a) > len(b) {
		return comp(a[len(b):], b)
	}

	if len(a) < len(b) {
		return comp(a, b[len(a):])
	}

	return false
}
