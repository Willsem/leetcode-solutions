import (
	"sort"
	"strconv"
	"strings"
)

func sortJumbled(mapping []int, nums []int) []int {
	m := make(map[int][]int)
	keys := make([]int, len(m))
	for _, n := range nums {
		ns := strconv.Itoa(n)
		var sb strings.Builder
		for _, d := range ns {
			di := int(d - '0')
			md := mapping[di]
			mds := strconv.Itoa(md)
			sb.WriteString(mds)
		}
		mappedString := sb.String()
		mappedNumber, _ := strconv.Atoi(mappedString)
		if m[mappedNumber] == nil {
			m[mappedNumber] = make([]int, 0)
			keys = append(keys, mappedNumber)
		}
		m[mappedNumber] = append(m[mappedNumber], n)
	}

	sort.Ints(keys)

	res := make([]int, 0)
	for _, v := range keys {
		res = append(res, m[v]...)
	}
	return res
}
