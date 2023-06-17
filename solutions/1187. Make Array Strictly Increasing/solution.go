import (
	"math"
	"sort"
)

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	dp := make(map[int]int, 1)
	dp[-1] = 0

	sort.Ints(arr2)

	for i := range arr1 {
		newDp := make(map[int]int, 0)

		for prev := range dp {
			if arr1[i] > prev {
				if _, ok := newDp[arr1[i]]; !ok {
					newDp[arr1[i]] = dp[prev]
				} else {
					newDp[arr1[i]] = min(newDp[arr1[i]], dp[prev])
				}
			}

			idx := binarySearch(0, len(arr2), func(i int) bool {
				return arr2[i] <= prev
			})

			if idx < len(arr2) {
				if _, ok := newDp[arr2[idx]]; !ok {
					newDp[arr2[idx]] = 1 + dp[prev]
				} else {
					newDp[arr2[idx]] = min(newDp[arr2[idx]], 1+dp[prev])
				}

			}
		}

		dp = newDp
	}

	ans := math.MaxInt
	for _, v := range dp {
		ans = min(ans, v)
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
