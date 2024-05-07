/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	head = reverse(head)

	curr := head
	var prev *ListNode
	next := 0
	for curr != nil {
		res := 2*curr.Val + next

		next = res / 10
		curr.Val = res % 10
		prev = curr
		curr = curr.Next
	}

	if next > 0 {
		prev.Next = &ListNode{next, nil}
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
