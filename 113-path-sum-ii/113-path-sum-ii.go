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
    inorderDfs(root,targetSum, []int{}, &result)
	return result
}
func inorderDfs(root *TreeNode, targetSum int, paths []int, result *[][]int) {
	if root == nil {
		return
	}
    // fmt.Println("Starting node: ",root.Val)
	targetSum -= root.Val
	paths = append(paths, root.Val)
    
    inorderDfs(root.Left, targetSum, paths, result)
    
    // fmt.Println("Back to: ", root.Val)
    // fmt.Println("Paths on this node are: ", paths)

    
    if targetSum == 0 && root.Left == nil && root.Right == nil {
        // fmt.Println("Adding the following path to result: ", paths)
        newPath := make([]int, len(paths))
        copy(newPath, paths)
		*result = append(*result, newPath)
	}
    
    inorderDfs(root.Right, targetSum, paths, result)
}