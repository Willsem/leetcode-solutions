func garbageCollection(garbage []string, travel []int) int {
	for i := 1; i < len(travel); i++ {
		travel[i] = travel[i-1] + travel[i]
	}

	garbageLastPos := make(map[rune]int)
	ans := 0
	for i := 0; i < len(garbage); i++ {
		for _, c := range garbage[i] {
			garbageLastPos[c] = i
		}
		ans += len(garbage[i])
	}

	garbageTypes := "MPG"
	for _, c := range garbageTypes {
		ans += func() int {
			if idx, ok := garbageLastPos[c]; ok && idx != 0 {
				return travel[idx-1]
			}
			return 0
		}()
	}

	return ans
}
