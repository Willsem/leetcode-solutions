const (
	firstWeek = 1 + 2 + 3 + 4 + 5 + 6 + 7
	step      = 7
)

func totalMoney(n int) int {
	weeks := n / 7
	days := n % 7

	forWeeks := int(float64(2*firstWeek+step*(weeks-1)) / 2.0 * float64(weeks))
	forDays := int(float64(2*(weeks+1)+days-1) / 2.0 * float64(days))

	return forWeeks + forDays
}
