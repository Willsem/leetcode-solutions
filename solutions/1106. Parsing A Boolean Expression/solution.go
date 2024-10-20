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

func parseBoolExpr(expression string) bool {
	s := NewStack[rune]()

	for _, v := range expression {
		switch v {
		case 't', 'f', '!', '&', '|':
			s.Push(v)

		case ')':
			hasTrue, hasFalse := false, false
			for s.Peek() != '!' && s.Peek() != '&' && s.Peek() != '|' {
				curr := s.Pop()
				if curr == 't' {
					hasTrue = true
				} else if curr == 'f' {
					hasFalse = true
				}
			}

			operator := s.Pop()
			switch operator {
			case '!':
				if hasTrue {
					s.Push(rune('f'))
				} else {
					s.Push(rune('t'))
				}
			case '&':
				if hasFalse {
					s.Push(rune('f'))
				} else {
					s.Push(rune('t'))
				}
			case '|':
				if hasTrue {
					s.Push(rune('t'))
				} else {
					s.Push(rune('f'))
				}
			}
		}
	}

	if s.Pop() == 't' {
		return true
	}

	return false
}
