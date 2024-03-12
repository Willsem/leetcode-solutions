/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	preHead := &ListNode{Next: head}
	curr := preHead

	prefixSum := 0
	prefixSumMap := make(map[int]*ListNode)
	prefixSumMap[0] = preHead

	for curr != nil {
		prefixSum += curr.Val
		prefixSumMap[prefixSum] = curr
		curr = curr.Next
	}

	prefixSum = 0
	curr = preHead

	for curr != nil {
		prefixSum += curr.Val
		curr.Next = prefixSumMap[prefixSum].Next
		curr = curr.Next
	}

	return preHead.Next
}
