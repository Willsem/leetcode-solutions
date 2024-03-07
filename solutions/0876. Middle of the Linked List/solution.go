/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	tempByOne, tempByTwo := head, head
	for tempByTwo != nil && tempByTwo.Next != nil {
		tempByOne = tempByOne.Next
		tempByTwo = tempByTwo.Next.Next
	}
	return tempByOne
}
