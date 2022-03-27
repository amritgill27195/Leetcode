/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    approach : using inorder 
    
    why inorder?
    - Because inorder goes Left -> Root (Process) -> Right
    - Which means, we will go down all the left first,
    - Hit a left.nil base case and come back to the left leaf
    - Then we can check if this node is a leaf node
    - if it is , we will save its path
    - how do we get its path?
    - At each node, we will main a string "path"
    - Once we hit leaf - we will add this string path to some output list
    

*/
func binaryTreePaths(root *TreeNode) []string {
    out := []string{}
    inorderDfs(root, "", &out)
    return out
}


func inorderDfs(root *TreeNode, paths string, out *[]string) {
    
    // base
    if root == nil {
        return
    }
    
    // logic
    paths += fmt.Sprintf("%v", root.Val)
    if root.Left != nil || root.Right != nil {
        paths += "->"
    }
    
    
    inorderDfs(root.Left, paths, out)
    if root.Left == nil && root.Right == nil {
        *out = append(*out, paths)
    }
    inorderDfs(root.Right, paths, out)
    
}