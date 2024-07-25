func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	qsort(nums, 0, len(nums))
	return nums
}

func qsort(nums []int, l, r int) {
	n := r - l
	if n <= 1 {
		return
	}

	pivot := choosePivot(nums[l], nums[l+n/2], nums[r-1])

	p, q := l, l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[i], nums[p] = nums[p], nums[i]
			if p == q {
				q++
			}
			p++
		}
		if nums[i] == pivot {
			nums[i], nums[q] = nums[q], nums[i]
			q++
		}
	}

	qsort(nums, l, p)
	qsort(nums, q, r)
}

func choosePivot(a, b, c int) int {
	if a >= b && a <= c || a <= b && a >= c {
		return a
	}
	if b >= a && b <= c || b <= a && b >= c {
		return b
	}
	return c
}
