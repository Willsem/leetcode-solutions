import "fmt"

func getLucky(s string, k int) int {
	num := 0
	for _, v := range s {
		num += sumOfDigits(int(v-'a') + 1)
	}

	k--
	for k > 0 {
		fmt.Println(num)
		num = sumOfDigits(num)
		k--
	}

	return num
}

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
