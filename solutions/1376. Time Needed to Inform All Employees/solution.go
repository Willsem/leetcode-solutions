import "fmt"

type point struct {
	id    int
	depth int
}

type Queue struct {
	data []point
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]point, 0),
	}
}

func (q *Queue) Clear() {
	q.data = []point{}
}

func (q *Queue) Push(el point) {
	q.data = append(q.data, el)
}

func (q *Queue) Peek() point {
	return q.data[0]
}

func (q *Queue) Pop() point {
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

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	graph := make([][]int, len(manager))
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := range manager {
		if manager[i] != -1 {
			graph[manager[i]] = append(graph[manager[i]], i)
		}
	}

	for _, v := range graph {
		fmt.Println(v)
	}

	result := 0
	visited := make([]bool, len(manager))

	q := NewQueue()
	q.Push(point{headID, 0})
	visited[headID] = true

	for !q.IsEmpty() {
		curr := q.Pop()

		if curr.depth > result {
			result = curr.depth
		}

		for _, v := range graph[curr.id] {
			if !visited[v] {
				visited[v] = true
				q.Push(point{v, curr.depth + informTime[curr.id]})
			}
		}
	}

	return result
}
