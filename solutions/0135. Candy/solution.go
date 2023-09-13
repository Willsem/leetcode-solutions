func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	ans, up, down, peak := 1, 0, 0, 0
	for i := 0; i < len(ratings)-1; i++ {
		prev, curr := ratings[i], ratings[i+1]

		if prev < curr {
			up++
			down = 0
			peak = up
			ans += 1 + up
		} else if prev == curr {
			up, down, peak = 0, 0, 0
			ans += 1
		} else {
			up = 0
			down++
			ans += 1 + down
			if peak >= down {
				ans--
			}
		}
	}

	return ans
}
