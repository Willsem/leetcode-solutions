import "strconv"

func hammingWeight(num uint32) int {
	count := 0
	for _, v := range strconv.FormatInt(int64(num), 2) {
		if v == '1' {
			count++
		}
	}
	return count
}
