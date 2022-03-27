/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, ans := postOrderDfs(root)
    return ans
}

func postOrderDfs(root *TreeNode) (int, bool) {
    
    // base
    if root == nil {
        return 0, true
    }
    
    
    lh,leftBalanced := postOrderDfs(root.Left)
    if !leftBalanced {
        return -1, false
    }
    
    rh, rightBalanced := postOrderDfs(root.Right)
    if !rightBalanced {
        return -1, false
    }
    
    diff := abs(lh-rh)
    if diff > 1 {
        return -1, false
    }
    max := lh
    if rh > max {max = rh} 
    return max+1, true
    
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}