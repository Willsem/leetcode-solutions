/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	tempByOne := head.Next
	tempByTwo := head.Next.Next

	for {
		if tempByOne == tempByTwo {
			return true
		}

		if tempByOne == nil || tempByTwo == nil || tempByTwo.Next == nil {
			return false
		}

		tempByOne = tempByOne.Next
		tempByTwo = tempByTwo.Next.Next
	}

	return false
}
