import "strconv"

func countSeniors(details []string) int {
	count := 0
	for _, pass := range details {
		ageString := pass[11:13]
		age, _ := strconv.Atoi(ageString)
		if age > 60 {
			count++
		}
	}
	return count
}
