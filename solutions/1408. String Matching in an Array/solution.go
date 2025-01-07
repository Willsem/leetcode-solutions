import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	resultSet := make(map[string]struct{})
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				resultSet[words[i]] = struct{}{}
				break
			}
		}
	}

	result := make([]string, 0, len(resultSet))
	for word := range resultSet {
		result = append(result, word)
	}
	return result
}
