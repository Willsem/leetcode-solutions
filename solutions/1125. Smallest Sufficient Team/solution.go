type Team struct {
	bits, count int
}

func (t Team) getPeopleList() []int {
	peopleList := []int{}
	mask := t.bits
	for i := 0; mask > 0; i++ {
		if mask%2 == 1 {
			peopleList = append(peopleList, i)
		}
		mask /= 2
	}
	return peopleList
}

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	sn, pn := len(req_skills), len(people)

	skillIdx := map[string]int{}
	for j, skill := range req_skills {
		skillIdx[skill] = j
	}

	peopleSkillMask := make([]int, pn)
	for i := range people {
		for _, skill := range people[i] {
			j := skillIdx[skill]
			peopleSkillMask[i] |= 1 << j
		}
	}

	dp := make([]Team, 1<<sn)

	dp[0] = Team{0, 0}
	for skillsMask := 1; skillsMask < len(dp); skillsMask++ {
		dp[skillsMask] = Team{(1 << pn) - 1, pn}
	}

	for skillsMask := 0; skillsMask < len(dp)-1; skillsMask++ {
		team := dp[skillsMask]

		for i := range people {
			if team.bits&(1<<i) == 1 {
				continue
			}

			if peopleSkillMask[i] == 0 {
				continue
			}

			newTeamCount := team.count + 1
			newSkillMask := skillsMask | peopleSkillMask[i]

			if newTeamCount < dp[newSkillMask].count {
				dp[newSkillMask].bits = team.bits | (1 << i)
				dp[newSkillMask].count = newTeamCount
			}
		}
	}

	allSkillsMask := 1<<sn - 1
	return dp[allSkillsMask].getPeopleList()
}
