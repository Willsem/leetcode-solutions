import "math/bits"

func longestSubsequence(s string, k int) int {
	sum := 0
	count := 0
	bits := bits.Len(uint(k))
	for i := 0; i < len(s); i++ {
		ch := s[len(s)-1-i]
		if ch == '1' {
			if i < bits && sum+(1<<i) <= k {
				sum += 1 << i
				count++
			}
		} else {
			count++
		}
	}
	return count
}
