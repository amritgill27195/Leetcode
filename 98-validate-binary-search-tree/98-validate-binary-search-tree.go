/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    Sorted order == valid BST
    How to traverse a tree in a sorted order?
    - Inorder traversal of a BST guarentees a sorted order
    
    If you traverse a tree in inorder fashion, you will see nodes in asc order.
    
    approach 1: Populate a list while traversing inorder
    - inorder traverse
        - time: o(n) - where n is the number of nodes in the tree. We end up seeing all nodes.
        - space: o(h) - where h is the height of the tree. Regardless if we use recursion or iterative stack
    - push each item to a list
        - time: o(1)
        - space: o(n)
    - finally check if the list is sorted
        - time : o(n)
        - space : o(1)
    time: o(n) + o(1) + o(n) = o(n)
    space: o(n) + o(h) + o(1) = o(n)
    
    
    approach 2: Maintain a prev pointer to compare current root node with
    - inorder traversal
    - compare prev with current root
        - prev will always be smaller in a true BST ( prev in a number line thats sorted will always be smalled )
        - if prev >= current { return false }
    time: o(n)
    space: o(h)
    
    approach 3: Iterative implmentation of approach #2
*/


// inorder DFS ( recursive ) -- using prev pointer
// time: o(n) - we visit all nodes in best/worst case
// space: o(h) - max height of the tree will be in recursion stack at worse case and o(n) in a skewed tree

type dfs struct {
    prev *TreeNode
}
func (d *dfs) inorderDfs(root *TreeNode) bool {
    
    // base
    if root == nil {return true }
    
    // logic
    leftValid := d.inorderDfs(root.Left)
    if !leftValid {
        return false
    }
    
    if d.prev != nil {
        if d.prev.Val >= root.Val {
            return false
        }
    }
    d.prev = root
    return d.inorderDfs(root.Right)
}


func isValidBST(root *TreeNode) bool {
    // d := &dfs{}
    // return d.inorderDfs(root)
    return preOrderUsingMinAndMax(root, nil, nil)
}




// cannot convert inorder to preorder as is, we have to define min and max boundaries at each node
// if we go left, min will be current min and max will be current root
// if we go right, min will be current root and max will be current max 
func preOrderUsingMinAndMax(root, min, max *TreeNode) bool {
    if root == nil {return true}
    if min != nil && root.Val <= min.Val {return false}
    if max != nil && root.Val >= max.Val {return false}
    
    return preOrderUsingMinAndMax(root.Left, min, root) && preOrderUsingMinAndMax(root.Right, root, max) 
}













