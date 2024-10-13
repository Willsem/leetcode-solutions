import "sort"

func smallestRange(nums [][]int) []int {
	ansx, ansy := -100000, 100000
	totalLists := len(nums)

	vp := make([][]int, 0)
	for i := 0; i < totalLists; i++ {
		for j := 0; j < len(nums[i]); j++ {
			vp = append(vp, []int{nums[i][j], i})
		}
	}

	sort.Slice(vp, func(i, j int) bool {
		return vp[i][0] < vp[j][0]
	})

	um := make(map[int]int)
	l, r := 0, 0
	n := len(vp)

	for r < n {
		list := vp[r][1]

		um[list]++

		for len(um) >= totalLists {
			if chk(vp[l][0], vp[r][0], ansx, ansy) {
				ansx = vp[l][0]
				ansy = vp[r][0]
			}

			oldList := vp[l][1]
			um[oldList]--

			if um[oldList] == 0 {
				delete(um, oldList)
			}
			l++
		}

		r++
	}

	return []int{ansx, ansy}
}

func chk(a, b, c, d int) bool {
	if b-a < d-c {
		return true
	} else if b-a == d-c && a < c {
		return true
	}
	return false
}
