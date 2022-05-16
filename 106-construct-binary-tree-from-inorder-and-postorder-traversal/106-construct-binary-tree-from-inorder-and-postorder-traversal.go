/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    postOrderIdx := len(postorder)-1
    
    inorderMap := map[int]int{}
    for idx, ele := range inorder {
        inorderMap[ele] = idx
    }
    
    var dfs func(start, end int) *TreeNode
    dfs = func(start, end int) *TreeNode{
        // base
        if start > end {
            return nil
        }
        
        // logic
        root := &TreeNode{Val: postorder[postOrderIdx]}
        postOrderIdx--
        rootIdxInInorder := inorderMap[root.Val]
        root.Right = dfs(rootIdxInInorder+1, end)
        root.Left = dfs(start, rootIdxInInorder-1)
        return root
    }
    
    return dfs(0, len(postorder)-1)
}