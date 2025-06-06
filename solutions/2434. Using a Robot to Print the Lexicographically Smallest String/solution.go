import "strings"

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

func robotWithString(s string) string {
	counts := make(map[byte]int)
	for i := range s {
		counts[s[i]]++
	}

	st := NewStack[byte]()
	result := &strings.Builder{}
	minChar := byte('a')

	for i := range s {
		c := s[i]

		st.Push(c)
		counts[c]--

		for minChar != 'z' && counts[minChar] == 0 {
			minChar++
		}

		for !st.IsEmpty() && st.Peek() <= minChar {
			result.WriteByte(st.Pop())
		}
	}

	return result.String()
}
