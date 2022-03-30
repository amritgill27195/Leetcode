/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    approach 1: level order using BFS
    This was the very first thing that came to my mind.
    Before processing each level ( in a queue ), we will save the last element from it
    and then process the nodes in queue like any other BFS

    time: o(n) - where n is the number of nodes in the tree. We visit and see every single node best/worst case ( we do not exit early anywhere )
    space: o(math.Max(widthOfTree)) or o(w) -- if its a skewed tree it becomes o(n)

    ------------------------------------------------

    approach 2: level order using DFS
    We can do level order traversal using DFS too by maintaining a level at each node in recursion stack.
    But how do we only save the right side of the tree?
    We will have a result array and each idx in this result array will represent a level in the tree
    
    2 ways:
    1. We can traverse left first and override each new element at the same level.
    2. We can traverse right first and only add to result array at the level IF the size of result array == level (i.e this is the first time we have seen this level )
    
    time: o(n) - where n is the number of nodes in the tree. We visit and see every single node best/worst case ( we do not exit early anywhere )
    space: o(h) - where h is the max height of tree and it will be o(n) space in a skewed tree
   
   
*/

// level order using BFS
func rightSideView(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }
    result := []int{}
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        qSize := len(queue)
        result = append(result, queue[qSize-1].Val)
        for qSize != 0 {
            dq := queue[0]
            queue = queue[1:]
            if dq.Left != nil { queue = append(queue, dq.Left )}
            if dq.Right != nil { queue = append(queue, dq.Right )}
            qSize--
        }
    }
    return result
}