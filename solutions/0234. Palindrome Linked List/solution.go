/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	temp := head
	fastPtr := head
	for fastPtr != nil && fastPtr.Next != nil {
		temp = temp.Next
		fastPtr = fastPtr.Next.Next
	}

	var prev *ListNode
	for temp != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}

	temp = prev
	for head != nil && temp != nil {
		if head.Val != temp.Val {
			return false
		}
		head = head.Next
		temp = temp.Next
	}

	return true
}
