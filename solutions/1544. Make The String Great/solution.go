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

func makeGood(s string) string {
	st := NewStack[byte]()
	st.Push(s[0])

	for i := 1; i < len(s); i++ {
		if !st.IsEmpty() && diff(st.Peek(), s[i]) == 32 {
			st.Pop()
			continue
		}
		st.Push(s[i])
	}
	return string(st.data)
}

func diff(a, b byte) byte {
	if b > a {
		return b - a
	}
	return a - b
}
