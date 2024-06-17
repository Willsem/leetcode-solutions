import "math"

func judgeSquareSum(c int) bool {
	for i := 0; i <= int(math.Sqrt(float64(c))); i++ {
		j := binarySearch(0, i+1, func(j int) bool {
			return i*i+j*j <= c
		}, false)

		if i*i+j*j == c {
			return true
		}
	}
	return false
}

func binarySearch(l, r int, comp func(i int) bool, leftSearch bool) int {
	for l < r-1 {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	if leftSearch {
		return r
	}

	return l
}
