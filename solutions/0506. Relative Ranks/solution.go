import (
	"sort"
	"strconv"
)

type scorePosition struct {
	score    int
	position int
}

func findRelativeRanks(score []int) []string {
	result := make([]string, len(score))
	sorted := make([]scorePosition, len(score))

	for i := range sorted {
		sorted[i] = scorePosition{
			score:    score[i],
			position: i,
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].score > sorted[j].score
	})

	result[sorted[0].position] = "Gold Medal"
	if len(sorted) > 1 {
		result[sorted[1].position] = "Silver Medal"
	}
	if len(sorted) > 2 {
		result[sorted[2].position] = "Bronze Medal"
	}

	for i := 3; i < len(sorted); i++ {
		result[sorted[i].position] = strconv.Itoa(i + 1)
	}

	return result
}
