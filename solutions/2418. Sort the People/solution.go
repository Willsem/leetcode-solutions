import "sort"

type person struct {
	name   string
	height int
}

func sortPeople(names []string, heights []int) []string {
	persons := make([]person, len(names))
	for i := range names {
		persons[i].name = names[i]
		persons[i].height = heights[i]
	}
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].height > persons[j].height
	})

	for i, person := range persons {
		names[i] = person.name
	}
	return names
}
