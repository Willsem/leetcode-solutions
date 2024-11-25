import (
	"fmt"
	"strings"
)

func findShortestPath(start, target string) int {
	if target == start {
		return 0
	}
	moves := [6][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	queue := make([]string, 0)
	nMoves := 0
	queue = append(queue, start)
	visited := map[string]bool{start: true}
	for len(queue) > 0 {
		length := len(queue)

		for length > 0 {
			current := queue[0]
			queue = queue[1:]

			if current == target {
				return nMoves
			}
			zeroPos := strings.IndexRune(current, '0')
			for _, next := range moves[zeroPos] {
				state := []byte(current)
				state[zeroPos], state[next] = state[next], state[zeroPos]
				nextState := string(state)
				if !visited[nextState] {
					visited[nextState] = true
					queue = append(queue, nextState)
				}
			}
			length--
		}

		nMoves++
	}
	return -1
}

func slidingPuzzle(board [][]int) int {
	target := "123450"
	start := ""

	for _, row := range board {
		for _, val := range row {
			start = start + fmt.Sprintf("%d", val)
		}
	}
	return findShortestPath(start, target)
}
