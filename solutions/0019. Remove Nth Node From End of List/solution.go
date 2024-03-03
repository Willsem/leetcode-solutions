/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	node := head
	for node != nil {
		node = node.Next
		length++
	}

	node = head
	var prev *ListNode
	for i := 0; i < length-n-1; i++ {
		prev = node
		node = node.Next
	}

	if prev == nil {
		return node.Next
	}

	prev.Next = node.Next
	return head
}
