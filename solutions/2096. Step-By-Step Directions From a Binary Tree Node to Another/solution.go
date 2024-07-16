import "strings"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type dest int

const (
	L dest = iota
	R
)

func getDirections(root *TreeNode, startValue int, destValue int) string {
	pathToStart := make([]dest, 0)
	pathToDest := make([]dest, 0)
	foundStart, foundDest := false, false

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || (foundStart && foundDest) {
			return
		}

		if node.Val == startValue {
			foundStart = true
		}
		if node.Val == destValue {
			foundDest = true
		}

		if !foundStart {
			pathToStart = append(pathToStart, L)
		}
		if !foundDest {
			pathToDest = append(pathToDest, L)
		}
		dfs(node.Left)
		if !foundStart {
			pathToStart = pathToStart[:len(pathToStart)-1]
		}
		if !foundDest {
			pathToDest = pathToDest[:len(pathToDest)-1]
		}

		if !foundStart {
			pathToStart = append(pathToStart, R)
		}
		if !foundDest {
			pathToDest = append(pathToDest, R)
		}
		dfs(node.Right)
		if !foundStart {
			pathToStart = pathToStart[:len(pathToStart)-1]
		}
		if !foundDest {
			pathToDest = pathToDest[:len(pathToDest)-1]
		}
	}

	dfs(root)

	startPathI := 0
	for startPathI < len(pathToStart) && startPathI < len(pathToDest) && pathToStart[startPathI] == pathToDest[startPathI] {
		startPathI++
	}

	res := strings.Builder{}
	for i := len(pathToStart) - 1; i >= startPathI; i-- {
		res.WriteString("U")
	}
	for i := startPathI; i < len(pathToDest); i++ {
		switch pathToDest[i] {
		case L:
			res.WriteString("L")
		case R:
			res.WriteString("R")
		}
	}

	return res.String()
}
