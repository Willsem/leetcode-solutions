type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(el int) {
	s.data = append(s.data, el)
}

func (s *Stack) Pop() int {
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()

	return s.Peek()
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func asteroidCollision(asteroids []int) []int {
	stack := NewStack()

	answer := make([]int, 0)
	for _, asteroid := range asteroids {
		if asteroid > 0 {
			stack.Push(asteroid)
		} else {
			addToAns := true
			for !stack.IsEmpty() {
				meet := stack.Pop()
				if meet > -asteroid {
					stack.Push(meet)
					break
				}

				if meet == -asteroid {
					addToAns = false
					break
				}
			}

			if stack.IsEmpty() && addToAns {
				answer = append(answer, asteroid)
			}
		}
	}

	return append(answer, stack.data...)
}
