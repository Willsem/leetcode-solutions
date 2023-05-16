/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil

	for curr != nil && curr.Next != nil {
		next := curr.Next

		if prev != nil {
			prev.Next = next
		} else {
			head = next
		}

		curr.Next = next.Next
		next.Next = curr

		prev = curr
		curr = curr.Next
	}

	return head
}
