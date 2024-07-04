/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	curr := head
	var lastZero *ListNode
	currSum := 0
	for curr != nil {
		if curr.Val == 0 {
			if currSum > 0 {
				lastZero.Next = &ListNode{
					Val:  currSum,
					Next: curr,
				}
				currSum = 0
			}

			lastZero = curr
		}

		currSum += curr.Val
		curr = curr.Next
	}

	head = head.Next
	curr = head
	var prev *ListNode
	for curr != nil {
		if curr.Val == 0 {
			prev.Next = curr.Next
		} else {
			prev = curr
		}

		curr = curr.Next
	}

	return head
}
