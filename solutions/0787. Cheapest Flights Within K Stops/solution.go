import "math"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Clear() {
	q.data = []T{}
}

func (q *Queue[T]) Push(el T) {
	q.data = append(q.data, el)
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) Pop() T {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

type Connection struct {
	city int
	cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	q := NewQueue[Connection]()
	q.Push(Connection{src, 0})

	minCost, maxStops := math.MaxInt64, k+1

	costOfFlights := make([][]int, n)
	for i := range costOfFlights {
		costOfFlights[i] = make([]int, n)
	}

	flightsMap := make(map[int][]int)
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]

		flightsMap[from] = append(flightsMap[from], to)
		costOfFlights[from][to] = cost
	}

	cityPathCost := make([]int, n)
	for i := range cityPathCost {
		cityPathCost[i] = math.MaxInt64
	}
	cityPathCost[src] = 0

	for !q.IsEmpty() && maxStops >= 0 {
		maxStops--
		size := q.Size()

		for size > 0 {
			size--
			c := q.Pop()

			if c.city == dst {
				if minCost > c.cost {
					minCost = c.cost
				}
				continue
			}

			for _, v := range flightsMap[c.city] {
				newCost := c.cost + costOfFlights[c.city][v]

				if cityPathCost[v] > newCost {
					q.Push(Connection{v, newCost})
					cityPathCost[v] = newCost
				}
			}
		}
	}

	if minCost == math.MaxInt64 {
		return -1
	}
	return minCost
}
