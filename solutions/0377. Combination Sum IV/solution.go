import "sort"

func combinationSum4Recursive(nums []int, target int, memo map[int]int) int {
	if val, found := memo[target]; found {
		return val
	}

	if target == 0 {
		return 1
	}

	if target < nums[0] {
		return 0
	}

	count := 0
	for _, num := range nums {
		if target-num < 0 {
			break
		}
		count += combinationSum4Recursive(nums, target-num, memo)
	}

	memo[target] = count
	return count
}

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	memo := make(map[int]int)
	return combinationSum4Recursive(nums, target, memo)
}
