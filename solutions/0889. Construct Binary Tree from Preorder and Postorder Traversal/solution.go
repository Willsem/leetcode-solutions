/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}

	leftRootVal := preorder[1]
	leftRootIndex := 0
	for i, v := range postorder {
		if v == leftRootVal {
			leftRootIndex = i
			break
		}
	}

	root.Left = constructFromPrePost(preorder[1:leftRootIndex+2], postorder[:leftRootIndex+1])
	root.Right = constructFromPrePost(preorder[leftRootIndex+2:], postorder[leftRootIndex+1:len(postorder)-1])

	return root
}
