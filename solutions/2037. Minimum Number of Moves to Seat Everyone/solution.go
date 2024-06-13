import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	res := 0
	for i := range seats {
		res += abs(seats[i] - students[i])
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
