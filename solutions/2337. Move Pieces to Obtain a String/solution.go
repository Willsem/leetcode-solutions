func canChange(start string, target string) bool {
	starti, targeti := 0, 0

	for starti < len(start) || targeti < len(target) {
		for starti < len(start) && start[starti] == '_' {
			starti++
		}

		for targeti < len(target) && target[targeti] == '_' {
			targeti++
		}

		if starti == len(start) || targeti == len(target) {
			return starti == len(start) && targeti == len(target)
		}

		if start[starti] != target[targeti] || (start[starti] == 'L' && starti < targeti) || (start[starti] == 'R' && starti > targeti) {
			return false
		}

		starti++
		targeti++
	}

	return true
}
