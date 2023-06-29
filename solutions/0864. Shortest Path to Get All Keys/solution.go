func shortestPathAllKeys(grid []string) int {
	m := len(grid)
	n := len(grid[0])

	keyBit := make(map[byte]int)
	bitStart := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isLowerCase(grid[i][j]) {
				keyBit[grid[i][j]] = bitStart
				bitStart++
			}
		}
	}

	maskEnd := (1 << bitStart) - 1
	maskSize := 1 << bitStart

	memo := make([][][]bool, m)
	for i := range memo {
		memo[i] = make([][]bool, n)
		for j := range memo[i] {
			memo[i][j] = make([]bool, maskSize)
		}
	}

	var start []int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' {
				start = []int{i, j, 0}
			}
		}
	}

	queue := [][]int{start}
	step := 0

	for len(queue) > 0 {
		sz := len(queue)

		for k := 0; k < sz; k++ {
			current := queue[0]
			queue = queue[1:]
			row := current[0]
			col := current[1]
			mask := current[2]

			if row < 0 || row >= m || col < 0 || col >= n {
				continue
			}

			if grid[row][col] == '#' {
				continue
			}

			if isUpperCase(grid[row][col]) {
				lowerKey := toLowerCase(grid[row][col])
				if (mask & (1 << keyBit[lowerKey])) == 0 {
					continue
				}
			}

			if isLowerCase(grid[row][col]) {
				mask |= 1 << keyBit[grid[row][col]]
			}

			if mask == maskEnd {
				return step
			}

			if memo[row][col][mask] {
				continue
			}

			memo[row][col][mask] = true

			queue = append(queue, []int{row + 1, col, mask})
			queue = append(queue, []int{row - 1, col, mask})
			queue = append(queue, []int{row, col + 1, mask})
			queue = append(queue, []int{row, col - 1, mask})
		}

		step++
	}

	return -1
}

func isLowerCase(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isUpperCase(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func toLowerCase(ch byte) byte {
	return ch + 32
}
