import (
	"math"
	"sort"
	"strconv"
	"strings"
)

const mod = 1e9 + 7

func countGoodIntegers(n int, k int) int64 {
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = (factorial[i-1] * i) % mod
	}

	digitMultisets := make(map[string]bool)

	min := int(math.Pow10(n - 1))
	max := int(math.Pow10(n)) - 1

	halfLen := (n + 1) / 2
	start := int(math.Pow10(halfLen - 1))
	end := int(math.Pow10(halfLen)) - 1

	for firstHalf := start; firstHalf <= end; firstHalf++ {
		s := strconv.Itoa(firstHalf)
		var palindromeStr string
		if n%2 == 0 {
			palindromeStr = s + reverseString(s)
		} else {
			palindromeStr = s + reverseString(s)[1:]
		}

		palindrome, _ := strconv.Atoi(palindromeStr)
		if palindrome < min || palindrome > max {
			continue
		}

		if palindrome%k == 0 {
			digits := strings.Split(palindromeStr, "")
			sort.Strings(digits)
			multisetKey := strings.Join(digits, "")
			digitMultisets[multisetKey] = true
		}
	}

	count := 0

	for multiset := range digitMultisets {
		digitCounts := make([]int, 10)
		for _, ch := range multiset {
			digitCounts[ch-'0']++
		}

		permutations := factorial[n]
		for _, cnt := range digitCounts {
			if cnt > 1 {
				permutations = permutations * modInverse(factorial[cnt], mod) % mod
			}
		}

		if digitCounts[0] > 0 {
			invalidPerms := factorial[n-1]
			for d, cnt := range digitCounts {
				if d == 0 {
					invalidPerms = invalidPerms * modInverse(factorial[cnt-1], mod) % mod
				} else {
					invalidPerms = invalidPerms * modInverse(factorial[cnt], mod) % mod
				}
			}
			permutations = (permutations - invalidPerms + mod) % mod
		}

		count = (count + permutations) % mod
	}

	return int64(count)
}

func modInverse(a int, mod int) int {
	return powMod(a, mod-2, mod)
}

func powMod(a, b, mod int) int {
	result := 1
	a = a % mod
	for b > 0 {
		if b%2 == 1 {
			result = (result * a) % mod
		}
		a = (a * a) % mod
		b = b / 2
	}
	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
