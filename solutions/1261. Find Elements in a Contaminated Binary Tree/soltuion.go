/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	elements map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	f := FindElements{
		elements: map[int]struct{}{},
	}

	if root == nil {
		return f
	}

	elements := make(map[int]struct{})

	var dfs func(node *TreeNode, parent int, isLeft bool)
	dfs = func(node *TreeNode, parent int, isLeft bool) {
		if node == nil {
			return
		}

		val := 2*parent + 1
		if !isLeft {
			val++
		}

		node.Val = val
		elements[val] = struct{}{}
		dfs(node.Left, val, true)
		dfs(node.Right, val, false)
	}

	root.Val = 0
	elements[0] = struct{}{}
	dfs(root.Left, 0, true)
	dfs(root.Right, 0, false)

	f.elements = elements
	return f
}

func (f *FindElements) Find(target int) bool {
	_, ok := f.elements[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
