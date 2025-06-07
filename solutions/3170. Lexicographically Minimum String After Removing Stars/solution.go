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

func clearStars(s string) string {
	stacks := make(map[byte]*Stack[int])
	sarr := []byte(s)

	for i, c := range sarr {
		if c != '*' {
			if stacks[c] == nil {
				stacks[c] = NewStack[int]()
			}
			stacks[c].Push(i)
		} else {
			for j := byte('a'); j <= 'z'; j++ {
				if stacks[j] != nil && !stacks[j].IsEmpty() {
					sarr[stacks[j].Pop()] = '*'
					break
				}
			}
		}
	}

	result := &strings.Builder{}
	for _, c := range sarr {
		if c != '*' {
			result.WriteByte(c)
		}
	}

	return result.String()
}
