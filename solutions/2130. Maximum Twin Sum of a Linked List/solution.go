/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	temp := head
	mid := head
	var prevMid *ListNode = nil
	for temp != nil {
		temp = temp.Next.Next
		prevMid = mid
		mid = mid.Next
	}

	prevMid.Next = nil

	prevMid = nil
	for mid != nil {
		temp = mid.Next
		mid.Next = prevMid
		prevMid = mid
		mid = temp
	}

	left := head
	right := prevMid
	max := left.Val + right.Val
	for left != nil {
		sum := left.Val + right.Val
		if sum > max {
			max = sum
		}
		left = left.Next
		right = right.Next
	}

	return max
}
