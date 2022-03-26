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
    inorderDfs(root,targetSum,nil, &result)
	return result
}
func inorderDfs(root *TreeNode, targetSum int, paths []int, result *[][]int) {
	if root == nil {
		return
	}
    
	targetSum -= root.Val
    paths = append(paths, root.Val)

    inorderDfs(root.Left, targetSum, paths, result)

    if targetSum == 0 && root.Left == nil && root.Right == nil {
        newPath := make([]int, len(paths))
        copy(newPath, paths)
		*result = append(*result, newPath)
	}
    
    inorderDfs(root.Right, targetSum, paths, result)

}