import "sort"

func fullBloomFlowers(flowers [][]int, person []int) []int {
	result := make([]int, len(person))
	begin, end := make([]int, 0, len(flowers)), make([]int, 0, len(flowers))
	for _, f := range flowers {
		begin = append(begin, f[0])
		end = append(end, f[1]+1)
	}
	sort.Ints(begin)
	sort.Ints(end)
	for i, day := range person {
		beginSize := binarySearchLE(day, begin)
		endSize := binarySearchLE(day, end)
		result[i] = beginSize - endSize
	}
	return result
}

func binarySearchLE(val int, array []int) int {
	l, r := 0, len(array)-1
	for l <= r {
		m := (l + r) >> 1
		if array[m] <= val {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
