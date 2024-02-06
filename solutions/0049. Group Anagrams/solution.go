import (
	"bytes"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		counts := make([]int, 26)
		for _, c := range str {
			counts[int(c-'a')]++
		}

		var strmap bytes.Buffer
		for i, count := range counts {
			strmap.WriteString(string('a'+byte(i)) + strconv.Itoa(count))
		}

		groups[strmap.String()] = append(groups[strmap.String()], str)
	}

	ans := make([][]string, 0, len(groups))
	for _, group := range groups {
		ans = append(ans, group)
	}
	return ans
}
