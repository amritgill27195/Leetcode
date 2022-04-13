/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

    If we are told that the two cousins x and y are somewhere closer to the top,
    then use BFS ( or ask this question )
    because BFS processes ALL nodes of a parent before proceeding to the next level ( breadth first )
    where as DFS processes 1 node at a time for a parent and may not visit the 2nd child after a very long time if the tree is huge.
    
    So ask - where are they potentially located or use this time to answer why BFS over DFS or vice versa.


    The general thing here is we are given x and y int values. We have to find whether these 2 are cousins.
    They can only be cousins if both conditions are true:
        1. On the same level ( hint hint, level order traversal )
        2. And their parents are not the same - i.e they both belong to diff parents
            - for this we likely need to pass down the parent of all the childs in the current level to compare x and y parents.
            - so we need to store/pass-down the node and the parent of this node.

    approach 1: level order using BFS
    - We can use a classic BFS and perform a level order traversal
    - Except in the queue, when we add a node, we will add a pair ( node and its parent )
    - Then when we are processing a node
        - if this node is x || y, save it and save its corresponding parent ( part of the node pair ) to xParent || yParent
    - however if we found x || y but not both after a level, there is no way these 2 can be cousins, so exit early.
    - once we have both x and y parent we can stop our level order traversal.
    - Finally compare xParent and yParent.
    
    time: o(n) - where n is the number of nodes 
    space: o(maxWidthOfTree) -- which will be the last level ( n/2 at max )

    
    approach 2: level order using DFS
    - Since we need do to a simple level order, we can use DFS as well. If levels are needed, we can maintain it at each node state
    - In this case levels are needed, because levels help us define whether x and y are on the same level.
    - The idea is we can do a simple DFS looking for both nodes.
    - As we are going down the tree we do need their each nodes parent, we will also save that as local state in each recursion call.
    - If current node == x || y, save its xLevel || yLevel and save its xParent || yParent ( whichever is found )
    - We can also exit the recursion early if xParent && yParent != nil ( i.e we found both, so exit )
    
    time: o(n) - where n is the number of nodes
    space: o(h) - where h is height of the tree, o(n) space in case of a skewed tree
*/

// dfs level order
func isCousins(root *TreeNode, x int, y int) bool {
    if root == nil {
        return false
    }
    
    var (
        xParent *TreeNode
        yParent *TreeNode
        xLevel int
        yLevel int
    )
    
    var dfs func(level int, a, parent *TreeNode)
    dfs = func(level int, a, parent *TreeNode) {
        // base
        if (xParent != nil && yParent != nil) || a == nil  {
            return
        }
        
        // logic
        if a.Val == x {
            xParent = parent
            xLevel = level
        }
        if a.Val == y {
            yParent = parent
            yLevel = level
        }
        
        dfs(level+1, a.Left, a)
        dfs(level+1, a.Right, a)
    }
    
    dfs(0, root, nil)
    if xLevel == yLevel {
        return xParent != yParent
    }
    
    return false
}






