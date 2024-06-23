import "container/list"

type Pair struct {
	Key   int
	Value int
}

func adjust(mono *list.List, left int) {
	for mono.Len() > 0 {
		elem := mono.Front().Value.(Pair)
		if elem.Value < left {
			mono.Remove(mono.Front())
		} else {
			break
		}
	}
}

func longestSubarray(nums []int, limit int) int {
	ans := 0
	monoMax := list.New()
	monoMin := list.New()

	j := 0
	for i := 0; i < len(nums); i++ {
		for monoMax.Len() > 0 {
			back := monoMax.Back().Value.(Pair)
			if nums[i] > back.Key {
				monoMax.Remove(monoMax.Back())
			} else {
				break
			}
		}
		for monoMin.Len() > 0 {
			back := monoMin.Back().Value.(Pair)
			if nums[i] < back.Key {
				monoMin.Remove(monoMin.Back())
			} else {
				break
			}
		}

		monoMin.PushBack(Pair{nums[i], i})
		monoMax.PushBack(Pair{nums[i], i})

		for {
			adjust(monoMin, j)
			adjust(monoMax, j)
			if monoMax.Front().Value.(Pair).Key-monoMin.Front().Value.(Pair).Key > limit {
				j++
			} else {
				break
			}
		}

		ans = max(ans, i-j+1)
	}

	return ans
}
