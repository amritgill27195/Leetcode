/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    Validate BST == inorder traversal
    
    which means, if we perform a inorder traversal on a valid BST,
    then nodes will be processed in a sorted manner
    
    So how do we verify that there sorted? Sure we can visit them but how to verify?
    
    appraoch 1:
    store each visited node in a list
    and compare the list
    
    time:
    space: o(h) + o(n)
    
    
    approach 2:
    Just like LinkedLists, maintain a prev pointer.
    Previous pointer here is the previous node for a current node in "inorder" fashion.
    Which means go all the left first
    Then set the prev
    Then use the prev to compare with next root node
    if its a valid bst, the prev must be smaller each time its compared with a next inorder root node
    ( which also means, if its valid, update the prev to current so the next inorder can compare its value with the correct prev one. )
    
    

*/
func isValidBST(root *TreeNode) bool {
	b := &bst{
		prev: nil,
		flag: true,
	}
	b.inorder(root)
	return b.flag
}


// the hack class to scope global var within an instance of a class..
// or if we expose global at main func level, than we have global pollution when multiple tests are ran
type bst struct {
	prev *TreeNode
	flag bool
}

func (b *bst) inorder(root *TreeNode) {
	if root == nil {
		return
	}
	b.inorder(root.Left)
	if b.prev != nil {
		if b.prev.Val >= root.Val {
			b.flag = false
		}
	}
	b.prev = root
	b.inorder(root.Right)
}
