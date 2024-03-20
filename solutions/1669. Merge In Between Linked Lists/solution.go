/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prev, temp := list1, list1
	n := 0
	for n < a {
		n++
		prev = temp
		temp = temp.Next
	}

	prev.Next = list2
	list2End := list2
	for list2End.Next != nil {
		list2End = list2End.Next
	}

	for n < b {
		n++
		temp = temp.Next
	}

	list2End.Next = temp.Next

	return list1
}
