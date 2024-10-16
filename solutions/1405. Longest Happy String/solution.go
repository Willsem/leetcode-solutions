import "strings"

func longestDiverseString(a int, b int, c int) string {
	result := strings.Builder{}
	currA, currB, currC := 0, 0, 0

	total := a + b + c
	for i := 0; i < total; i++ {
		if (a >= b && a >= c && currA != 2) || (a > 0 && (currB == 2 || currC == 2)) {
			result.WriteString("a")
			a--
			currA++
			currB, currC = 0, 0
		} else if (b >= a && b >= c && currB != 2) || (b > 0 && (currC == 2 || currA == 2)) {
			result.WriteString("b")
			b--
			currB++
			currA, currC = 0, 0
		} else if (c >= a && c >= b && currC != 2) || (c > 0 && (currA == 2 || currB == 2)) {
			result.WriteString("c")
			c--
			currC++
			currA, currB = 0, 0
		}
	}

	return result.String()
}
