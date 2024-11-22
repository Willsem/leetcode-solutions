import "strings"

func maxEqualRowsAfterFlips(matrix [][]int) int {
	patternCounter := make(map[string]int)

	for _, row := range matrix {
		pattern := &strings.Builder{}
		for _, col := range row {
			if col == row[0] {
				pattern.WriteString("t")
			} else {
				pattern.WriteString("f")
			}
		}
		patternCounter[pattern.String()]++
	}

	maxFreq := 0
	for _, freq := range patternCounter {
		maxFreq = max(maxFreq, freq)
	}
	return maxFreq
}
