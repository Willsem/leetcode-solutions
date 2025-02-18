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

func smallestNumber(pattern string) string {
	s := NewStack[int]()
	result := make([]byte, 0, len(pattern)+1)

	for i := 0; i <= len(pattern); i++ {
		s.Push(i + 1)

		if i == len(pattern) || pattern[i] == 'I' {
			for !s.IsEmpty() {
				result = append(result, byte(s.Pop())+'0')
			}
		}
	}

	return string(result)
}
