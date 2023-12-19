import "math"

var dist = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 0},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := range img {
		res[i] = make([]int, len(img[i]))
		for j := range img[i] {
			res[i][j] = calcSmooth(img, i, j)
		}
	}
	return res
}

func calcSmooth(img [][]int, i, j int) int {
	sum, count := 0, 0
	for _, d := range dist {
		i := i + d[0]
		j := j + d[1]

		if i < 0 || i >= len(img) {
			continue
		}

		if j < 0 || j >= len(img[i]) {
			continue
		}

		sum += img[i][j]
		count++
	}

	return int(math.Floor(float64(sum) / float64(count)))
}
