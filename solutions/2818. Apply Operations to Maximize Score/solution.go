import "sort"

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

func maximumScore(nums []int, k int) int {
	n := len(nums)

	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	primeScores := make([]int, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		if primeScores[i] == 0 {
			for j := i; j <= maxNum; j += i {
				primeScores[j]++
			}
		}
	}

	nextGreater := make([]int, n)
	prevGreater := make([]int, n)

	s := NewStack[int]()
	for i := 0; i < n; i++ {
		for !s.IsEmpty() && primeScores[nums[s.Peek()]] < primeScores[nums[i]] {
			nextGreater[s.Pop()] = i
		}
		s.Push(i)
	}
	for !s.IsEmpty() {
		nextGreater[s.Pop()] = n
	}

	for i := n - 1; i >= 0; i-- {
		for !s.IsEmpty() && primeScores[nums[s.Peek()]] <= primeScores[nums[i]] {
			prevGreater[s.Pop()] = i
		}
		s.Push(i)
	}
	for !s.IsEmpty() {
		prevGreater[s.Pop()] = -1
	}

	type Element struct {
		val   int
		count int
	}

	elements := make([]Element, n)
	for i := 0; i < n; i++ {
		left := prevGreater[i]
		right := nextGreater[i]
		count := (i - left) * (right - i)
		elements[i] = Element{nums[i], count}
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].val > elements[j].val
	})

	score := 1
	remaining := k

	for _, elem := range elements {
		if remaining <= 0 {
			break
		}
		take := elem.count
		if take > remaining {
			take = remaining
		}
		score = (score * pow(elem.val, take)) % mod
		remaining -= take
	}

	return score
}

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		n /= 2
	}
	return res
}
