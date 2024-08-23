import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {
	nums := make([][]int, 0)

	currNum := make([]int, 0)
	currStr := ""
	for _, r := range expression {
		switch r {
		case '/':
			num, _ := strconv.Atoi(currStr)
			currNum = append(currNum, num)
			currStr = ""

		case '+', '-':
			if len(currStr) > 0 {
				num, _ := strconv.Atoi(currStr)
				currNum = append(currNum, num)
				nums = append(nums, currNum)
				currNum = make([]int, 0)
			}
			currStr = string(r)

		default:
			currStr += string(r)
		}
	}

	num, _ := strconv.Atoi(currStr)
	currNum = append(currNum, num)
	nums = append(nums, currNum)

	resNum, resDenom := 0, 1
	for _, num := range nums {
		resNum = resNum*num[1] + num[0]*resDenom
		resDenom *= num[1]
	}

	gcd := abs(findGCD(resNum, resDenom))
	resNum /= gcd
	resDenom /= gcd

	return fmt.Sprintf("%d/%d", resNum, resDenom)
}

func findGCD(a, b int) int {
	if a == 0 {
		return b
	}
	return findGCD(b%a, a)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
