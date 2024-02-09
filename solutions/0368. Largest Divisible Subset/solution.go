import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	groupSize := make([]int, len(nums))
	prevEl := make([]int, len(nums))
	maxIndex := 0

	for i := range nums {
		groupSize[i] = 1
		prevEl[i] = -1

		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && groupSize[j]+1 > groupSize[i] {
				groupSize[i] = groupSize[j] + 1
				prevEl[i] = j
			}
		}

		if groupSize[i] > groupSize[maxIndex] {
			maxIndex = i
		}
	}

	res := make([]int, 0)
	for i := maxIndex; i != -1; i = prevEl[i] {
		res = append([]int{nums[i]}, res...)
	}
	return res
}
