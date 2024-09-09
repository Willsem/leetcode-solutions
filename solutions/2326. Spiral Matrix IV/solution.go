/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
const (
	RIGHT = iota
	DOWN
	LEFT
	UP
	COUNT
)

var directions = [][]int{
	{0, +1}, // RIGHT
	{+1, 0}, // DOWN
	{0, -1}, // LEFT
	{-1, 0}, // UP
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	curr := head
	i, j := 0, 0
	dir := RIGHT
	for curr != nil {
		matrix[i][j] = curr.Val

		nextI, nextJ := i+directions[dir][0], j+directions[dir][1]
		if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || matrix[nextI][nextJ] != -1 {
			dir = (dir + 1) % COUNT
			nextI, nextJ = i+directions[dir][0], j+directions[dir][1]
		}

		i, j = nextI, nextJ
		curr = curr.Next
	}

	return matrix
}
