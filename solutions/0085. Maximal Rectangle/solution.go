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

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0])+1)
	heights[len(heights)-1] = -1
	res := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		s := NewStack[int]()
		for i := range heights {
			for !s.IsEmpty() && heights[s.Peek()] > heights[i] {
				prev := heights[s.Pop()]
				width := i
				if !s.IsEmpty() {
					width = i - s.Peek() - 1
				}
				res = max(res, prev*width)
			}
			s.Push(i)
		}
	}

	return res
}
