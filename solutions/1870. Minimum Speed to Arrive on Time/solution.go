import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	return binarySearch(1, 1e7, func(speed int) bool {
		time := 0.0
		for i := range dist {
			t := float64(dist[i]) / float64(speed)
			if i == len(dist)-1 {
				time += t
			} else {
				time += math.Ceil(t)
			}
		}
		return time <= hour
	})
}

func binarySearch(l, r int, comp func(i int) bool) int {
	res := -1

	for l <= r {
		mid := (l + r) / 2
		if comp(mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}
