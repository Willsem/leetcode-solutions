func findTheDifference(s string, t string) byte {
	var ans byte
	for i := range s {
		ans ^= s[i]
		ans ^= t[i]
	}
	return ans ^ t[len(t)-1]
}
