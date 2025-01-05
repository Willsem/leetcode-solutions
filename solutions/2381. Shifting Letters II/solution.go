func shiftingLetters(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)

	for _, sh := range shifts {
		start, end, dir := sh[0], sh[1], sh[2]
		switch dir {
		case 1:
			diff[start]++
			diff[end+1]--
		case 0:
			diff[start]--
			diff[end+1]++
		}
	}

	currShift := 0
	result := make([]byte, len(s))
	for i := range s {
		currShift += diff[i]
		shifted := (int(s[i]-'a') + currShift) % 26
		if shifted < 0 {
			shifted += 26
		}
		result[i] = byte(shifted) + 'a'
	}

	return string(result)
}
