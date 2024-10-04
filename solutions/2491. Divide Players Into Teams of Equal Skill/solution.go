func dividePlayers(skill []int) int64 {
	total := 0
	skillMap := make(map[int]int)
	for _, s := range skill {
		total += s
		skillMap[s]++
	}

	half := len(skill) / 2
	if total%half != 0 {
		return -1
	}

	target := total / half
	var result int64 = 0

	for currSkill, currFreq := range skillMap {
		partner := target - currSkill
		if partnerCount, ok := skillMap[partner]; !ok || currFreq != partnerCount {
			return -1
		}

		result += int64(currSkill) * int64(partner) * int64(currFreq)
	}

	return result / 2
}
