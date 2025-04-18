import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return rle(countAndSay(n - 1))
}

func rle(in string) string {
	prev, count := 0, 0
	res := &strings.Builder{}
	for i := range in {
		num, _ := strconv.Atoi(string(in[i]))
		if count == 0 || num == prev {
			count++
		} else {
			res.WriteString(fmt.Sprintf("%d%d", count, prev))
			count = 1
		}

		prev = num
	}
	res.WriteString(fmt.Sprintf("%d%d", count, prev))
	return res.String()
}
