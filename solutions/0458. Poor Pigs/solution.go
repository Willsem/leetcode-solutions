import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	testRound := minutesToTest / minutesToDie
	testGroup := testRound + 1
	res := int(math.Ceil(math.Log2(float64(buckets)) / math.Log2(float64(testGroup))))
	return res
}
