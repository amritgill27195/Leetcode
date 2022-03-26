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
	inorderDfs(root,targetSum, 0, nil, &result)
	return result
}
func inorderDfs(root *TreeNode, targetSum, cs int, paths []int, result *[][]int) {
	if root == nil {
		return
	}
	cs += root.Val
	paths = append(paths, root.Val)

	inorderDfs(root.Left, targetSum, cs, paths, result)

    if cs == targetSum && root.Left == nil && root.Right == nil {
		*result = append(*result, append([]int{}, paths...))
	}


    
    inorderDfs(root.Right, targetSum, cs, paths, result)
}