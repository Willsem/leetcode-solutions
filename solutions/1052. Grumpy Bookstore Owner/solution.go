func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	unrealizedCustomers := 0
	for i := 0; i < minutes; i++ {
		unrealizedCustomers += customers[i] * grumpy[i]
	}

	maxUnrealizedCustomers := unrealizedCustomers
	for i := minutes; i < n; i++ {
		unrealizedCustomers += customers[i] * grumpy[i]
		unrealizedCustomers -= customers[i-minutes] * grumpy[i-minutes]

		if unrealizedCustomers > maxUnrealizedCustomers {
			maxUnrealizedCustomers = unrealizedCustomers
		}
	}

	totalCustomers := maxUnrealizedCustomers
	for i := 0; i < n; i++ {
		totalCustomers += customers[i] * (1 - grumpy[i])
	}

	return totalCustomers
}
