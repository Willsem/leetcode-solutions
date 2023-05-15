/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	n := 0
	for it := head; it != nil; it = it.Next {
		n++
	}

	i := k - 1
	nodeI := head

	j := n - k
	nodeJ := head

	for i > 0 || j > 0 {
		if i > 0 {
			i--
			nodeI = nodeI.Next
		}

		if j > 0 {
			j--
			nodeJ = nodeJ.Next
		}
	}

	nodeI.Val, nodeJ.Val = nodeJ.Val, nodeI.Val
	return head
}
