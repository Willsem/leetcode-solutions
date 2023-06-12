import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result, start, prev := make([]string, 0), nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev+1 {
			prev = nums[i]
		} else {
			result = updateResult(result, start, prev)
			start, prev = nums[i], nums[i]
		}
	}

	result = updateResult(result, start, prev)
	return result
}

func updateResult(result []string, start, prev int) []string {
	if start == prev {
		return append(result, strconv.Itoa(start))
	}

	return append(result, fmt.Sprintf("%d->%d", start, prev))
}
