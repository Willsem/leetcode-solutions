/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	return dfs(root, head, head)
}

func dfs(treeNode *TreeNode, listNode *ListNode, head *ListNode) bool {
	if listNode == nil {
		return true
	}

	if treeNode == nil {
		return false
	}

	var nextListNode *ListNode
	if treeNode.Val == listNode.Val {
		nextListNode = listNode.Next
	} else if treeNode.Val == head.Val {
		nextListNode = listNode
		head = head.Next
	} else {
		nextListNode = head
	}

	return dfs(treeNode.Left, nextListNode, head) || dfs(treeNode.Right, nextListNode, head)
}
