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

type Element struct {
	node  *TreeNode
	level int
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	q := NewQueue[Element]()
	q.Push(Element{root, 0})

	prevValue, currLevel := -1, -1

	for !q.IsEmpty() {
		curr := q.Pop()

		if curr.node == nil {
			continue
		}

		if curr.level > currLevel {
			prevValue = curr.node.Val
			currLevel = curr.level
			if currLevel%2 == 0 {
				if curr.node.Val%2 == 0 {
					return false
				}
			} else {
				if curr.node.Val%2 == 1 {
					return false
				}
			}
		} else {
			if currLevel%2 == 0 {
				if curr.node.Val%2 == 0 || prevValue >= curr.node.Val {
					return false
				}
			} else {
				if curr.node.Val%2 == 1 || prevValue <= curr.node.Val {
					return false
				}
			}
			prevValue = curr.node.Val
		}

		q.Push(Element{curr.node.Left, curr.level + 1})
		q.Push(Element{curr.node.Right, curr.level + 1})
	}

	return true
}
