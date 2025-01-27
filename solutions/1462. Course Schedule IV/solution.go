func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	depMatrix := make([][]bool, numCourses)
	for i := range depMatrix {
		depMatrix[i] = make([]bool, numCourses)
	}

	for _, prerequisite := range prerequisites {
		depMatrix[prerequisite[0]][prerequisite[1]] = true
	}

	for k := range numCourses {
		for i := range numCourses {
			for j := range numCourses {
				if depMatrix[i][k] && depMatrix[k][j] {
					depMatrix[i][j] = true
				}
			}
		}
	}

	results := make([]bool, len(queries))
	for i, query := range queries {
		results[i] = depMatrix[query[0]][query[1]]
	}

	return results
}
