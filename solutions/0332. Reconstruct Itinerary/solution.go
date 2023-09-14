import "sort"

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)

	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	for key := range graph {
		sort.Sort(sort.Reverse(sort.StringSlice(graph[key])))
	}

	stack := []string{"JFK"}
	var itinerary []string

	for len(stack) > 0 {
		for len(graph[stack[len(stack)-1]]) > 0 {
			last := len(graph[stack[len(stack)-1]]) - 1
			stack = append(stack, graph[stack[len(stack)-1]][last])
			graph[stack[len(stack)-2]] = graph[stack[len(stack)-2]][:last]
		}
		itinerary = append(itinerary, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	for i := 0; i < len(itinerary)/2; i++ {
		itinerary[i], itinerary[len(itinerary)-1-i] = itinerary[len(itinerary)-1-i], itinerary[i]
	}

	return itinerary
}
