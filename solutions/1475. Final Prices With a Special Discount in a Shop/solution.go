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

func finalPrices(prices []int) []int {
	s := NewStack[int]()

	result := make([]int, len(prices))
	for i := range prices {
		for !s.IsEmpty() && prices[s.Peek()] >= prices[i] {
			j := s.Pop()
			result[j] = prices[j] - prices[i]
		}
		s.Push(i)
	}

	for !s.IsEmpty() {
		i := s.Pop()
		result[i] = prices[i]
	}

	return result
}
