func numTeams(rating []int) int {
	n := len(rating)
	teams := 0

	for mid := 0; mid < n; mid++ {
		left, right := 0, 0

		for l := mid - 1; l >= 0; l-- {
			if rating[l] < rating[mid] {
				left++
			}
		}

		for r := mid + 1; r < n; r++ {
			if rating[r] > rating[mid] {
				right++
			}
		}

		teams += left * right

		left = mid - left
		right = n - mid - 1 - right

		teams += left * right
	}

	return teams
}
