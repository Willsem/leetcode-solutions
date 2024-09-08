func lenList(head *ListNode) int {
	n := 0
	for temp := head; temp != nil; temp = temp.Next {
		n++
	}
	return n
}
