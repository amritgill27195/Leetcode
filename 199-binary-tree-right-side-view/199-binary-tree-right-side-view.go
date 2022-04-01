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
    level == idx and if len(res) == level this means the max idx avail right now is len(res)-1 and if level == len(res) and level is our idx
    then that means we do not have level as an idx avail ( its less then level ) - so append in this case.
    
    2 ways:
    1. We can traverse left first and go reckless override each new element at the same level. 
        - if right goes after left, the last overrides will be the right side view
    2. We can traverse right first and only add to result array at the level IF the size of result array == level (i.e this is the first time we have seen this level)
        - gather all the right side values for each level
        - if this level does not exist in our result array
            - that means this is the first time we have seen this level, so append it with current level value
        - if this level does exist, it means we already have the farthest right side value for this level
    
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


// level order using DFS ( i.e maintain a level local state at each node )
// this exists to scope down global to only this class/struct or else global pollution happens
// type dfs struct {
//     result []int
// }
// func rightSideView(root *TreeNode) []int {
//     d := &dfs{result: []int{}}
//     d.dfsRight(root, 0)
//     return d.result
// }


// func (d *dfs) dfsRight(root *TreeNode, level int) {
//     if root == nil {
//         return
//     }
    
//     // logic
//     // if this level does not exist, add it
//     // else skip because we already have the far most right side for this level
//     if len(d.result) == level {
//         d.result = append(d.result, root.Val)
//     }
//     d.dfsRight(root.Right, level+1)
//     d.dfsRight(root.Left, level+1)
// }

// func (d *dfs) dfsLeft(root *TreeNode, level int) {
//     if root == nil {
//         return
//     }
    
//     // logic
//     if len(d.result) == level {
//         d.result = append(d.result, root.Val)
//     } else {
//         d.result[level] = root.Val
//     }
//     d.dfsLeft(root.Left, level+1)
//     d.dfsLeft(root.Right, level+1)
// }
