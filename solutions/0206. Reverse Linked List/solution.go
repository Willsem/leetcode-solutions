/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var ptr1 *ListNode
	ptr2 := head

	for ptr2 != nil {
		tmp := ptr2.Next
		ptr2.Next = ptr1
		ptr1, ptr2 = ptr2, tmp
	}

	return ptr1
}
