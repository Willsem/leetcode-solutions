import "math"

func maxMatrixSum(matrix [][]int) int64 {
	sum := 0
	minAbsValue := math.MaxInt
	countMinus := 0

	for _, row := range matrix {
		for _, v := range row {
			absV := int(math.Abs(float64(v)))
			sum += absV
			if v < 0 {
				countMinus++
			}
			minAbsValue = min(minAbsValue, absV)
		}
	}

	if countMinus%2 != 0 {
		sum -= 2 * minAbsValue
	}

	return int64(sum)
}
