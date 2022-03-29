/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    b := BSTIterator{stack: []*TreeNode{}}
    b.inorderDfs(root)
    return b
}

func (this *BSTIterator) inorderDfs(root *TreeNode) {
    for root != nil {
        this.stack = append(this.stack, root)
        root = root.Left
    }
}


func (this *BSTIterator) Next() int {
    if this.HasNext() {
        top := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        this.inorderDfs(top.Right)
        return top.Val
    }
    return -1
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */