/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	temp := head

	for temp != nil {
		n++
		temp = temp.Next
	}

	div := n / k
	mod := n % k

	res := make([]*ListNode, k)
	prev := head
	temp = head
	for i := 0; i < k; i++ {
		res[i] = temp

		if temp != nil {
			for j := 0; j < div; j++ {
				prev = temp
				temp = temp.Next
			}

			if mod > 0 {
				mod--
				prev = temp
				temp = temp.Next
			}

			prev.Next = nil
		}
	}

	return res
}
