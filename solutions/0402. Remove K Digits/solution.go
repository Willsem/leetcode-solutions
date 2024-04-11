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

func removeKdigits(num string, k int) string {
	s := NewStack[rune]()

	for _, digit := range num {
		for !s.IsEmpty() && k > 0 && s.Peek() > digit {
			s.Pop()
			k--
		}
		s.Push(digit)
	}

	for k > 0 && !s.IsEmpty() {
		s.Pop()
		k--
	}

	sb := strings.Builder{}
	for _, digit := range s.data {
		sb.WriteRune(digit)
	}

	res := sb.String()
	fmt.Println(res)
	i := 0
	for ; i < len(res) && res[i] == '0'; i++ {
	}

	if i >= len(res) {
		return "0"
	}

	return res[i:]
}
