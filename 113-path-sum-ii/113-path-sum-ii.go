/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	dfs(root,targetSum, 0, nil, &result)
	return result
}
func dfs(root *TreeNode, targetSum, cs int, paths []int, result *[][]int) {
	if root == nil {
		return
	}
	cs += root.Val
	paths = append(paths, root.Val)

	dfs(root.Left, targetSum, cs, paths, result)

    if cs == targetSum && root.Left == nil && root.Right == nil {
		*result = append(*result, append([]int{}, paths...))
	}


    
    dfs(root.Right, targetSum, cs, paths, result)
}