func buyChoco(prices []int, money int) int {
	min, secondMin := 101, 101
	for _, price := range prices {
		if price < min {
			min, secondMin = price, min
		} else if price < secondMin {
			secondMin = price
		}
	}

	if min+secondMin > money {
		return money
	}

	return money - min - secondMin
}
