import "sort"

func minOperations(nums []int) int {
	res := len(nums)

	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	newNums := make([]int, 0, len(set))
	for num := range set {
		newNums = append(newNums, num)
	}

	sort.Ints(newNums)

	j := 0
	for i := range newNums {
		for j < len(newNums) && newNums[j] < newNums[i]+len(nums) {
			j++
		}

		count := j - i
		if res > len(nums)-count {
			res = len(nums) - count
		}
	}

	return res
}
