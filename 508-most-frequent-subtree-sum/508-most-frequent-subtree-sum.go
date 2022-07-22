/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    freqMap := map[int]int{}
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {
            return math.MinInt64
        }
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        sum := r.Val
        if left != math.MinInt64 {
            freqMap[left]++
            sum += left
        }
        if right != math.MinInt64 {
            freqMap[right]++
            sum += right
        }
        return sum
    }
    rootSum := dfs(root)
    freqMap[rootSum]++
    maxV := math.MinInt64
    for _, v := range freqMap {
        maxV = max(v, maxV)
    }
    
    out := []int{}
    for k, v := range freqMap {
        if v == maxV {
            out = append(out, k)
        }
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}