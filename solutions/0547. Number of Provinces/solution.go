type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(el int) {
	s.data = append(s.data, el)
}

func (s *Stack) Pop() int {
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()

	return s.Peek()
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	s := NewStack()
	count := 0

	for i := range isConnected {
		if !visited[i] {
			count++
			visited[i] = true
			s.Push(i)

			for !s.IsEmpty() {
				curr := s.Pop()
				for j := range isConnected {
					if !visited[j] && isConnected[curr][j] == 1 {
						s.Push(j)
						visited[j] = true
					}
				}
			}
		}
	}

	return count
}
