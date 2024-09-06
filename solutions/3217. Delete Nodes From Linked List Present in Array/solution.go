/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		if _, ok := set[curr.Val]; !ok {
			prev = curr
			curr = curr.Next
			continue
		}

		if prev == nil {
			head = curr.Next
			curr = head
		} else {
			prev.Next = curr.Next
			curr = curr.Next
		}
	}

	return head
}
