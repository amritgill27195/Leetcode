/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    min := ""
    var dfs func(r *TreeNode, path string)
    dfs = func(r *TreeNode, path string) {
        // base
        if r == nil {
            return
        }
        //logic

        path += string('a'+r.Val)

        dfs(r.Left, path)
        if r.Left == nil && r.Right == nil {
            //// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
            // strings.Compare(a,b)
            revPath := reverse(path)
            if min == "" {
                min = revPath
            } else if val := strings.Compare(min, revPath); val == 1 {
                min = revPath
            }
        }
        
        dfs(r.Right, path)
    }
    dfs(root, "")
    return min
}

func reverse(s string) string {
    newS := new(strings.Builder)
    for i := len(s)-1; i >= 0; i-- {
        newS.WriteByte(s[i])
    }
    return newS.String()
} 