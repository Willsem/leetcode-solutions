/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	criticalCoords := make([]int, 0)
	prevValue := head.Val
	curr := head.Next
	for i := 0; ; i++ {
		if curr.Next == nil {
			break
		}

		if curr.Val > prevValue && curr.Val > curr.Next.Val || curr.Val < prevValue && curr.Val < curr.Next.Val {
			criticalCoords = append(criticalCoords, i)
		}

		prevValue = curr.Val
		curr = curr.Next
	}

	if len(criticalCoords) < 2 {
		return []int{-1, -1}
	}

	max := criticalCoords[len(criticalCoords)-1] - criticalCoords[0]
	min := max
	for i := 0; i < len(criticalCoords)-1; i++ {
		diff := criticalCoords[i+1] - criticalCoords[i]
		if min > diff {
			min = diff
		}
	}

	return []int{min, max}
}
