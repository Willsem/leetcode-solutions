import "sort"

type frequent struct {
	number   int
	frequent int
}

func topKFrequent(nums []int, k int) []int {
	frequentcyIndex := make(map[int]int, 0)
	frequentcy := make([]frequent, 0)
	for _, v := range nums {
		if _, ok := frequentcyIndex[v]; !ok {
			frequentcy = append(frequentcy, frequent{
				number:   v,
				frequent: 0,
			})
			frequentcyIndex[v] = len(frequentcy) - 1
		}
		frequentcy[frequentcyIndex[v]].frequent++
	}

	sort.Slice(frequentcy, func(i, j int) bool {
		return frequentcy[i].frequent > frequentcy[j].frequent
	})

	res := make([]int, 0)
	for i := 0; len(res) < k; i++ {
		res = append(res, frequentcy[i].number)
	}

	return res
}
