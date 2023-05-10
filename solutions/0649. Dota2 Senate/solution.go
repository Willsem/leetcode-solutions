func predictPartyVictory(senate string) string {
	banned := make([]bool, len(senate))

	radiantBannedCounter := 0
	direBannedCounter := 0

	for {
		// Count of senators which wasn't banned.
		radiantCount := 0
		direCount := 0

		for i, v := range senate {
			if banned[i] {
				continue
			}

			if v == 'R' {
				if radiantBannedCounter == 0 {
					// This senator could ban one nearest Dire senator.
					direBannedCounter++
					radiantCount++
				} else {
					// This senator couldn't make a choise because of ban.
					radiantBannedCounter--
					banned[i] = true
				}
			} else if v == 'D' {
				if direBannedCounter == 0 {
					// This senator could ban one nearest Radiant senator.
					radiantBannedCounter++
					direCount++
				} else {
					// This senator couldn't make a choise because of ban.
					direBannedCounter--
					banned[i] = true
				}
			}
		}

		if radiantCount == 0 {
			return "Dire"
		}

		if direCount == 0 {
			return "Radiant"
		}
	}
}
