var dir = map[rune][]int{
	'N': {0, 1},
	'S': {0, -1},
	'E': {1, 0},
	'W': {-1, 0},
}

func isPathCrossing(path string) bool {
	x, y := 0, 0
	points := make(map[int]map[int]struct{})
	points[x] = make(map[int]struct{})
	points[x][y] = struct{}{}

	for _, v := range path {
		x += dir[v][0]
		y += dir[v][1]

		if _, ok := points[x]; !ok {
			points[x] = make(map[int]struct{})
			points[x][y] = struct{}{}
			continue
		}

		if _, ok := points[x][y]; !ok {
			points[x][y] = struct{}{}
			continue
		}

		return true
	}

	return false
}
