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

// level order using dfs
type dfs struct {
    xParent, yParent *TreeNode
    xLevel, yLevel int
}
func isCousins(root *TreeNode, x int, y int) bool {
    d := &dfs{}
    d.levelOrderDfs(root, nil, 0, x, y)
    if d.xLevel == d.yLevel { // found cousins on the same level
        return d.xParent != d.yParent
    }
    return false
}

func (d *dfs)levelOrderDfs(root, parent *TreeNode, level, x, y int) {
    // base
    if root == nil {
        return 
    }
    
    
    
    d.levelOrderDfs(root.Left, root, level+1, x,y)
    
    // logic
    if root.Val == x {
        d.xParent = parent
        d.xLevel = level
    } else if root.Val == y {
        d.yParent = parent
        d.yLevel = level
    }
    if d.xParent != nil && d.yParent != nil {
        // exit early since we found both parents
        return
    }
    
    d.levelOrderDfs(root.Right, root, level+1, x,y)
}


// type pair struct {
//     node *TreeNode
//     parent *TreeNode
// }
// // level order using BFS
// func isCousins(root *TreeNode, x int, y int) bool {
    
//     if root == nil {
//         return false
//     }
    
//     var xParent *TreeNode
//     var yParent *TreeNode
//     rootPair := &pair{node: root, parent: nil}
//     queue := []*pair{rootPair}
    
//     for len(queue) != 0 {
//         qSize := len(queue)
//         for qSize != 0 {
//             dq := queue[0]
//             queue = queue[1:]
            
//             if dq.node.Val == x {
//                 xParent = dq.parent
//             } else if dq.node.Val == y {
//                 yParent = dq.parent
//             }
            
//             if dq.node.Left != nil {queue = append(queue, &pair{node: dq.node.Left, parent: dq.node })}
//             if dq.node.Right != nil {queue = append(queue, &pair{node: dq.node.Right, parent: dq.node })}
            
//             qSize--
//         }
//         // once we are done with a level and we only found parent for 1, exit early, there is no way these 2 are cousins.
//         if (xParent == nil && yParent != nil) ||  (xParent != nil && yParent == nil ) {
//             return false
//         }
//     }
    
//     // if we end up here, that means we found parents for both x and y from the same level 
//     return xParent != yParent
// }