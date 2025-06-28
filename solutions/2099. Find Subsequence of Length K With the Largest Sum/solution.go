import "sort"

type value struct {
	num int
	i   int
}

func maxSubsequence(nums []int, k int) []int {
	values := make([]value, len(nums))
	for i, num := range nums {
		values[i] = value{
			num: num,
			i:   i,
		}
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].num > values[j].num
	})

	values = values[:k]
	sort.Slice(values, func(i, j int) bool {
		return values[i].i < values[j].i
	})

	result := make([]int, len(values))
	for i := range values {
		result[i] = values[i].num
	}

	return result
}
