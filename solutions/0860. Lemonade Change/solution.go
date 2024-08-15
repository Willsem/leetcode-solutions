func lemonadeChange(bills []int) bool {
	count5, count10 := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			count5++
		case 10:
			count5--
			if count5 < 0 {
				return false
			}
			count10++
		case 20:
			if count10 == 0 {
				count5 -= 3
			} else {
				count5--
				count10--
			}
			if count5 < 0 {
				return false
			}
		}
	}

	return true
}
