func convertToTitle(columnNumber int) string {
	res := ""

	for columnNumber > 0 {
		columnNumber--
		res += string('A' + columnNumber%26)
		columnNumber /= 26
	}

	return reverse(res)
}

func reverse(s string) string {
	rns := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		rns[i], rns[len(s)-i-1] = rns[len(s)-i-1], rns[i]
	}
	return string(rns)
}
