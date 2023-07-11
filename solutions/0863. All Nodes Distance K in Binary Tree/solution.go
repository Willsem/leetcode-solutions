type data struct {
	node  *TreeNode
	index int
}

type Stack struct {
	data []data
}

func NewStack() *Stack {
	return &Stack{
		data: make([]data, 0),
	}
}

func (s *Stack) Push(el data) {
	s.data = append(s.data, el)
}

func (s *Stack) Pop() data {
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()

	return s.Peek()
}

func (s *Stack) Peek() data {
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

type depthData struct {
	index int
	depth int
}

type Queue struct {
	data []depthData
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]depthData, 0),
	}
}

func (q *Queue) Clear() {
	q.data = []depthData{}
}

func (q *Queue) Push(el depthData) {
	q.data = append(q.data, el)
}

func (q *Queue) Peek() depthData {
	return q.data[0]
}

func (q *Queue) Pop() depthData {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
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
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if root == nil {
		return []int{}
	}

	values := make([]int, 0)
	graph := make([][]int, 0)

	stack := NewStack()
	stack.Push(data{
		node:  root,
		index: 0,
	})

	values = append(values, root.Val)
	graph = append(graph, make([]int, 0))

	targetIndex := -1
	for !stack.IsEmpty() {
		d := stack.Pop()
		currNode := d.node

		if currNode == target {
			targetIndex = d.index
		}

		if currNode.Left != nil {
			index := len(values)
			stack.Push(data{
				node:  currNode.Left,
				index: index,
			})

			values = append(values, currNode.Left.Val)

			graph = append(graph, []int{d.index})
			graph[d.index] = append(graph[d.index], index)
		}

		if currNode.Right != nil {
			index := len(values)
			stack.Push(data{
				node:  currNode.Right,
				index: index,
			})

			values = append(values, currNode.Right.Val)

			graph = append(graph, []int{d.index})
			graph[d.index] = append(graph[d.index], index)
		}
	}

	q := NewQueue()
	q.Push(depthData{
		index: targetIndex,
		depth: 0,
	})
	result := make([]int, 0)
	visited := make([]bool, len(values))
	visited[targetIndex] = true

	for !q.IsEmpty() {
		curr := q.Pop()

		if curr.depth > k {
			return result
		}

		if curr.depth == k {
			result = append(result, values[curr.index])
			continue
		}

		for _, v := range graph[curr.index] {
			if !visited[v] {
				visited[v] = true
				q.Push(depthData{
					index: v,
					depth: curr.depth + 1,
				})
			}
		}
	}

	return result
}
