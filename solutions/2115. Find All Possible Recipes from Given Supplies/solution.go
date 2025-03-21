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

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	availableSupplies := make(map[string]struct{})
	recipeToIndex := make(map[string]int)
	dependencyGraph := make(map[string][]string)

	for _, supply := range supplies {
		availableSupplies[supply] = struct{}{}
	}

	for i := range recipes {
		recipeToIndex[recipes[i]] = i
	}

	inDegree := make([]int, len(recipes))

	for i := range recipes {
		for _, ingredient := range ingredients[i] {
			if _, ok := availableSupplies[ingredient]; !ok {
				dependencyGraph[ingredient] = append(dependencyGraph[ingredient], recipes[i])
				inDegree[i]++
			}
		}
	}

	q := NewQueue[int]()
	for i := range recipes {
		if inDegree[i] == 0 {
			q.Push(i)
		}
	}

	result := make([]string, 0)
	for !q.IsEmpty() {
		i := q.Pop()
		recipe := recipes[i]
		result = append(result, recipe)

		for _, dependentRecipe := range dependencyGraph[recipe] {
			inDegree[recipeToIndex[dependentRecipe]]--
			if inDegree[recipeToIndex[dependentRecipe]] == 0 {
				q.Push(recipeToIndex[dependentRecipe])
			}
		}
	}

	return result
}
