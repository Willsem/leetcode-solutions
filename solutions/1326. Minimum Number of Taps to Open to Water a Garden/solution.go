func minTaps(n int, ranges []int) int {
	arr := make([]int, n+1)

	for i, r := range ranges {
		if r == 0 {
			continue
		}
		left := max(0, i-r)
		arr[left] = max(arr[left], i+r)
	}

	end, far_can_reach, cnt := 0, 0, 0
	for i := 0; i <= n; i++ {
		if i > end {
			if far_can_reach <= end {
				return -1
			}
			end = far_can_reach
			cnt++
		}
		far_can_reach = max(far_can_reach, arr[i])
	}

	if end < n {
		cnt++
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
