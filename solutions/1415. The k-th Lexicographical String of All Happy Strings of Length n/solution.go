import "strings"

func getHappyString(n int, k int) string {
	if maxK(n) < k {
		return ""
	}

	chars := "abc"
	result := strings.Builder{}
	cnt := 0
	prev := len(chars)

	for i := 1; i <= n; i++ {
		p := 1 << (n - i)
		q := (k - cnt - 1) / p
		cnt += p * q
		if prev <= q {
			q++
		}

		result.WriteByte(chars[q])
		prev = q
	}

	return result.String()
}

func maxK(n int) int {
	return 3 * (1 << (n - 1))
}
