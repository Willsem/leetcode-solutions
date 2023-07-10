type node struct {
	node  *TreeNode
	depth int
}

type Queue struct {
	data []node
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]node, 0),
	}
}

func (q *Queue) Clear() {
	q.data = []node{}
}

func (q *Queue) Push(el node) {
	q.data = append(q.data, el)
}

func (q *Queue) Peek() node {
	return q.data[0]
}

func (q *Queue) Pop() node {
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := NewQueue()
	q.Push(node{
		node:  root,
		depth: 1,
	})

	for {
		el := q.Pop()

		if el.node.Left == nil && el.node.Right == nil {
			return el.depth
		}

		if el.node.Left != nil {
			q.Push(node{
				node:  el.node.Left,
				depth: el.depth + 1,
			})
		}

		if el.node.Right != nil {
			q.Push(node{
				node:  el.node.Right,
				depth: el.depth + 1,
			})
		}
	}
}
