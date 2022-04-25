/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    BST == inorder gives us the sorted form ( as long as the tree is valid bst )
    In this, we do have an invalid BST and we have to fix it by swapping **values**
    
    Everytime we have binaryTree/BST question, evaluate whether you have 2 different types of cases
    where 2 could be part of the same subtree or 2 nodes could be spread out in different subtrees
    
    Here we are given that there will be EXACTLY 2 nodes who are in bad state;
    which means;
        1. The 2 nodes could be in the same subtree
        2. The 2 nodes could be spread out into 2 different subtrees
    
    approach: brute force
    - use inorder to generate a list of node values
    - sort them
    - use the sorted list to write back into the tree using inorder

*/
func recoverTree(root *TreeNode)  {
    nodes := []int{}
    var listGenerator func(r *TreeNode)
    listGenerator = func(r *TreeNode) {
        if r == nil {return}
        listGenerator(r.Left)
        nodes = append(nodes, r.Val)
        listGenerator(r.Right)
    }
    listGenerator(root)
    
    sort.Ints(nodes)
    ptr := 0
    var inorderWriter func(r *TreeNode)
    inorderWriter = func(r *TreeNode) {
        if r == nil {
            return
        }
        inorderWriter(r.Left)
        r.Val = nodes[ptr]
        ptr++
        inorderWriter(r.Right)
    }
    inorderWriter(root)
}