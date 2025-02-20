import "strings"

func findDifferentBinaryString(nums []string) string {
	result := &strings.Builder{}
	for i := 0; i < len(nums); i++ {
		curr := nums[i][i]
		if curr == '0' {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}

	return result.String()
}
