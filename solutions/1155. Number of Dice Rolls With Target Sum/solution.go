func numRollsToTarget(dice int, faces int, target int) int {
	const mod = int(1e9 + 7)

	next, prev := make([]int, target+1), make([]int, target+1)

	for k := 1; k <= faces && k <= target; k++ {
		prev[k] = 1
	}

	for n := 1; n < dice; n++ {
		for t, cnt := range prev {
			if cnt == 0 {
				continue
			}

			for k := 1; k <= faces; k++ {
				if t+k > target {
					break
				}
				next[t+k] = (next[t+k] + cnt) % mod
			}
		}
		next, prev = prev, next
		clear(next)
	}
	return prev[target]
}
