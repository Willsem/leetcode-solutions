import "fmt"

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func robotSim(commands []int, obstacles [][]int) int {
	dirIndex := 0

	obstacleSet := make(map[string]struct{})
	for _, obstacle := range obstacles {
		key := getKey(obstacle[0], obstacle[1])
		obstacleSet[key] = struct{}{}
	}

	x, y := 0, 0
	maxDistance := 0

	for _, cmd := range commands {
		switch cmd {
		case -2:
			dirIndex = (dirIndex + 3) % 4
		case -1:
			dirIndex = (dirIndex + 1) % 4
		default:
			for step := 0; step < cmd; step++ {
				newX := x + directions[dirIndex][0]
				newY := y + directions[dirIndex][1]

				key := getKey(newX, newY)
				if _, exists := obstacleSet[key]; exists {
					break
				}
				x, y = newX, newY
				maxDistance = max(maxDistance, x*x+y*y)
			}
		}
	}

	return maxDistance
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
