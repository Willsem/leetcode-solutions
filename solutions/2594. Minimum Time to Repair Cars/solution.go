import "math"

func repairCars(ranks []int, cars int) int64 {
	result := binarySearch(0, cars*cars*ranks[0], func(mid int) bool {
		carsRepaired := 0
		for _, rank := range ranks {
			carsRepaired += int(math.Sqrt(float64(mid) / float64(rank)))
		}
		return carsRepaired < cars
	}, true)
	return int64(result)
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
