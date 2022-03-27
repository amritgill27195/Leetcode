/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var out []int
    dfs(root, &out)
    return out
}

func dfs(root *TreeNode, ans *[]int) {
    if root == nil {
        return
    }
    dfs(root.Left, ans)
    dfs(root.Right, ans)
    *ans = append(*ans, root.Val)
}