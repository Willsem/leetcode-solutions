/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for curr, next := head, head.Next; next != nil; curr, next = next, next.Next {
		mid := gcd(curr.Val, next.Val)

		curr.Next = &ListNode{
			Val:  mid,
			Next: next,
		}
	}

	return head
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
