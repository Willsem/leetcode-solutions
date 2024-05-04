import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	count := 0
	left, right := 0, len(people)-1
	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		count++
		right--
	}

	return count
}
