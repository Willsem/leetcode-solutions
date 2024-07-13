import "sort"

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

type Robot struct {
	Index     int
	Position  int
	Health    int
	Direction byte
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	robots := []*Robot{}
	for i := 0; i < len(positions); i++ {
		robot := &Robot{
			Index:     i,
			Position:  positions[i],
			Health:    healths[i],
			Direction: directions[i],
		}
		robots = append(robots, robot)
	}

	sort.SliceStable(robots, func(i, j int) bool {
		return robots[i].Position < robots[j].Position
	})

	stack := NewStack[*Robot]()
	for _, robot := range robots {
		for !stack.IsEmpty() {
			top := stack.Peek()
			if top.Direction == 'R' && robot.Direction == 'L' {
				stack.Pop()

				if top.Health == robot.Health {
					robot = nil
					break
				}

				if top.Health > robot.Health {
					top.Health--
					robot = top
				} else {
					robot.Health--
				}
			} else {
				break
			}
		}
		if robot != nil {
			stack.Push(robot)
		}
	}

	sort.SliceStable(stack.data, func(i, j int) bool {
		return stack.data[i].Index < stack.data[j].Index
	})

	results := []int{}
	for _, robot := range stack.data {
		results = append(results, robot.Health)
	}

	return results
}
