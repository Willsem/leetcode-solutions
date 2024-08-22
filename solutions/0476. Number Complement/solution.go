import "math"

func findComplement(num int) int {
	resNum := 0
	power := 0

	for num > 0 {
		bit := 0
		if num%2 == 0 {
			bit = 1
		}

		if power == 0 {
			resNum += bit
		} else {
			resNum += int(math.Pow(float64(bit*2), float64(power)))
		}

		power++
		num /= 2
	}

	return resNum
}
