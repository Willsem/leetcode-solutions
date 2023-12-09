type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(el T) {
	s.data = append(s.data, el)
}

func (s *Stack[T]) Pop() T {
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()

	return s.Peek()
}

func (s *Stack[T]) Peek() T {
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

type stackData struct {
	node *TreeNode
	was  bool
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	q := NewStack[stackData]()
	q.Push(stackData{
		node: root,
		was:  false,
	})

	for !q.IsEmpty() {
		data := q.Pop()
		node := data.node
		was := data.was
		if node == nil {
			continue
		}

		if was {
			res = append(res, node.Val)
			q.Push(stackData{
				node: node.Right,
				was:  false,
			})
		} else {
			q.Push(stackData{
				node: node,
				was:  true,
			})
			q.Push(stackData{
				node: node.Left,
				was:  false,
			})
		}
	}

	return res
}
