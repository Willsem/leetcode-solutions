const mod = 1e9 + 7

type Solution struct {
	table [][]int
}

func NewSolution(table [][]int) *Solution {
	return &Solution{
		table: table,
	}
}

func (s *Solution) dfs(nums []int) int {
	if len(nums) < 3 {
		return 1
	}

	leftNodes := make([]int, 0)
	rightNodes := make([]int, 0)
	for _, v := range nums {
		if v < nums[0] {
			leftNodes = append(leftNodes, v)
		} else if v > nums[0] {
			rightNodes = append(rightNodes, v)
		}
	}

	leftCount := s.dfs(leftNodes) % mod
	rightCount := s.dfs(rightNodes) % mod
	counts := (leftCount * rightCount) % mod

	return (counts * s.table[len(nums)-1][len(leftNodes)]) % mod
}

func numOfWays(nums []int) int {
	table := make([][]int, len(nums)+1)
	for i := range table {
		table[i] = make([]int, i+1)
		table[i][0] = 1
		table[i][i] = 1
		for j := 1; j < i; j++ {
			table[i][j] = (table[i-1][j-1] + table[i-1][j]) % mod
		}
	}

	s := NewSolution(table)
	return (s.dfs(nums) - 1) % mod
}
