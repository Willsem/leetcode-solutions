import (
	"strconv"
	"strings"
)

func minFlips(a int, b int, c int) int {
	binA := strconv.FormatInt(int64(a), 2)
	binB := strconv.FormatInt(int64(b), 2)
	binC := strconv.FormatInt(int64(c), 2)

	maxLen := len(binA)
	if len(binB) > maxLen {
		maxLen = len(binB)
	}
	if len(binC) > maxLen {
		maxLen = len(binC)
	}

	binA = strings.Repeat("0", maxLen-len(binA)) + binA
	binB = strings.Repeat("0", maxLen-len(binB)) + binB
	binC = strings.Repeat("0", maxLen-len(binC)) + binC

	count := 0
	for i := range binA {
		if binC[i] == '0' {
			if binA[i] == '1' {
				count++
			}
			if binB[i] == '1' {
				count++
			}
		} else {
			if binA[i] == '0' && binB[i] == '0' {
				count++
			}
		}
	}

	return count
}
