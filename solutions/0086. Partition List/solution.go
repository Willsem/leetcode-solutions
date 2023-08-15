/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var headBefore, currBefore *ListNode
	var headAfter, currAfter *ListNode

	temp := head
	for temp != nil {
		if temp.Val < x {
			if headBefore == nil {
				headBefore = temp
				currBefore = headBefore
			} else {
				currBefore.Next = temp
				currBefore = temp
			}
		} else {
			if headAfter == nil {
				headAfter = temp
				currAfter = headAfter
			} else {
				currAfter.Next = temp
				currAfter = temp
			}
		}
		temp = temp.Next
	}

	if headBefore == nil {
		return headAfter
	}

	currBefore.Next = headAfter

	if currAfter != nil {
		currAfter.Next = nil
	}

	return headBefore
}
