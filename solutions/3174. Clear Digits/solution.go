func clearDigits(s string) string {
	result := make([]byte, 0)
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' && len(result) > 0 {
			result = result[:len(result)-1]
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
