func minRemoveToMakeValid(s string) string {
	count := 0
	sarr := []byte(s)
	for i, v := range sarr {
		if v == '(' {
			count++
		} else if v == ')' {
			if count == 0 {
				sarr[i] = '*'
			} else {
				count--
			}
		}
	}

	for i := len(sarr) - 1; i >= 0 && count > 0; i-- {
		if sarr[i] == '(' {
			sarr[i] = '*'
			count--
		}
	}

	sb := strings.Builder{}
	for _, c := range sarr {
		if c != '*' {
			sb.WriteByte(c)
		}
	}
	return sb.String()
}
