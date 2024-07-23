import "sort"

func frequencySort(nums []int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	countsToNums := make([][]int, len(nums)+1)
	for num, count := range counts {
		countsToNums[count] = append(countsToNums[count], num)
	}

	res := make([]int, 0, len(nums))
	for count, nums := range countsToNums {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		for _, num := range nums {
			for i := 0; i < count; i++ {
				res = append(res, num)
			}
		}
	}

	return res
}
