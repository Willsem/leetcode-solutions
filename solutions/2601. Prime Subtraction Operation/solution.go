import "math"

func primeSubOperation(nums []int) bool {
	maxElement := nums[0]
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}

	sieve := make([]bool, maxElement+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[1] = false

	for i := 2; i <= int(math.Sqrt(float64(maxElement+1))); i++ {
		if sieve[i] {
			for j := i * i; j <= maxElement; j += i {
				sieve[j] = false
			}
		}
	}

	currValue := 1
	i := 0
	for i < len(nums) {
		difference := nums[i] - currValue

		if difference < 0 {
			return false
		}

		if sieve[difference] == true || difference == 0 {
			i++
			currValue++
		} else {
			currValue++
		}
	}
	return true
}
