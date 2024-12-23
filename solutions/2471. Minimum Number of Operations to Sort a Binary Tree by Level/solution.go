import "sort"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Clear() {
	q.data = []T{}
}

func (q *Queue[T]) Push(el T) {
	q.data = append(q.data, el)
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) Pop() T {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
	q := NewQueue[*TreeNode]()
	q.Push(root)
	swaps := 0

	for !q.IsEmpty() {
		size := q.Size()
		levelValues := make([]int, 0, size)

		for range size {
			node := q.Pop()
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}

		swaps += getMinSwaps(levelValues)
	}

	return swaps
}

func getMinSwaps(level []int) int {
	swaps := 0
	sorted := make([]int, len(level))
	copy(sorted, level)
	sort.Ints(sorted)

	positions := make(map[int]int, len(level))
	for i, v := range level {
		positions[v] = i
	}

	for i := range level {
		if level[i] != sorted[i] {
			swaps++

			currPos := positions[sorted[i]]
			positions[level[i]] = currPos
			level[currPos] = level[i]
		}
	}

	return swaps
}
