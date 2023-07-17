/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	result := &ListNode{}
	temp := result
	prevAddition := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + prevAddition

		temp.Val = sum % 10
		prevAddition = sum / 10

		l1 = l1.Next
		l2 = l2.Next

		if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}

	for l1 != nil {
		sum := l1.Val + prevAddition
		temp.Val = sum % 10
		prevAddition = sum / 10
		l1 = l1.Next

		if l1 != nil {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}

	for l2 != nil {
		sum := l2.Val + prevAddition
		temp.Val = sum % 10
		prevAddition = sum / 10
		l2 = l2.Next

		if l2 != nil {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}

	if prevAddition != 0 {
		temp.Next = &ListNode{}
		temp.Next.Val = prevAddition
	}

	return reverseList(result)
}

func reverseList(head *ListNode) *ListNode {
	temp := head
	var prev *ListNode
	for temp != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return prev
}
