const mod = 1e9 + 7

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

func sumSubarrayMins(arr []int) int {
	s := NewStack[int]()
	res := 0

	for j := 0; j < len(arr); j++ {
		for !s.IsEmpty() && arr[j] < arr[s.Peek()] {
			mid := s.Pop()
			left := -1
			if s.Size() > 0 {
				left = s.Peek()
			}
			right := j
			count := (mid - left) * (right - mid) % mod
			res += (arr[mid] * count) % mod
			res %= mod
		}

		s.Push(j)
	}

	right := len(arr)
	for !s.IsEmpty() {
		mid := s.Pop()
		left := -1
		if s.Size() > 0 {
			left = s.Peek()
		}
		count := (mid - left) * (right - mid) % mod
		res += (arr[mid] * count) % mod
		res %= mod
	}

	return res
}
