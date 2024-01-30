import "strconv"

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

func evalRPN(tokens []string) int {
	s := NewStack[int]()

	for _, token := range tokens {
		if tokenIsOperation(token) {
			num2, num1 := s.Pop(), s.Pop()
			s.Push(calc(num1, num2, token))
		} else {
			num, _ := strconv.Atoi(token)
			s.Push(num)
		}
	}

	return s.Peek()
}

func tokenIsOperation(token string) bool {
	if token == "+" || token == "-" || token == "*" || token == "/" {
		return true
	}
	return false
}

func calc(num1, num2 int, operation string) int {
	switch operation {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}

	return 0
}
