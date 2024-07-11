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

func reverseParentheses(s string) string {
	st := NewStack[int]()
	pairs := make(map[int]int)

	for i := range s {
		if s[i] == '(' {
			st.Push(i)
		} else if s[i] == ')' {
			leftI := st.Pop()
			pairs[leftI] = i
			pairs[i] = leftI
		}
	}

	res := strings.Builder{}
	direction := 1
	for i := 0; i < len(s); i += direction {
		if s[i] == '(' || s[i] == ')' {
			i = pairs[i]
			direction = -direction
		} else {
			res.WriteByte(s[i])
		}
	}

	return res.String()
}
