/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)

	curr := head
	var prev, prevPrev *ListNode

	maxVal := head.Val
	for curr != nil {
		if curr.Val >= maxVal {
			maxVal = curr.Val
		} else {
			prev.Next = curr.Next
			prev, curr = prevPrev, prev
		}

		prevPrev, prev, curr = prev, curr, curr.Next
	}

	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}
