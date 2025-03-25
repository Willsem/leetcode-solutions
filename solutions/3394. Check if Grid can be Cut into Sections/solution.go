import "sort"

const (
	SX = iota
	SY
	FX
	FY
)

func checkValidCuts(n int, rectangles [][]int) bool {
	if checkValidCutsByAxis(rectangles, SX, FX) {
		return true
	}
	return checkValidCutsByAxis(rectangles, SY, FY)
}

func checkValidCutsByAxis(rectangles [][]int, S, F int) bool {
	sort.Slice(rectangles, func(i, j int) bool {
		rectI, rectJ := rectangles[i], rectangles[j]
		if rectI[S] == rectJ[S] {
			return rectI[F] < rectJ[F]
		}

		return rectI[S] < rectJ[S]
	})

	countMerged := 1
	finish := rectangles[0][F]
	for _, rect := range rectangles[1:] {
		if finish > rect[S] {
			finish = max(finish, rect[F])
			continue
		}

		countMerged++
		finish = rect[F]
	}

	return countMerged >= 3
}
