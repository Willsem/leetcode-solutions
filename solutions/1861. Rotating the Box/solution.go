const (
	STONE    = '#'
	OBSTACLE = '*'
	EMPTY    = '.'
)

func rotateTheBox(box [][]byte) [][]byte {
	newBox := make([][]byte, len(box[0]))
	for i := range newBox {
		newBox[i] = make([]byte, len(box))
	}

	for i := range box {
		for j := range box[i] {
			newBox[j][len(newBox[j])-i-1] = box[i][j]
		}
	}

	for j := range newBox[0] {
		prevObstacle := -1
		for i := 0; i < len(newBox); i++ {
			stoneCount := 0

		LOOP:
			for ; i < len(newBox); i++ {
				switch newBox[i][j] {
				case STONE:
					stoneCount++
				case OBSTACLE:
					break LOOP
				}
			}

			for ii := i - 1; ii > prevObstacle; ii-- {
				if stoneCount > 0 {
					newBox[ii][j] = STONE
					stoneCount--
				} else {
					newBox[ii][j] = EMPTY
				}
			}

			prevObstacle = i
		}
	}

	return newBox
}
