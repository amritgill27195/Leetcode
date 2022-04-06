func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    
    var dfs func (c []int, i int, t int, p []int)
    dfs = func(c []int,i int, t int, p []int) {
        // base
        if t == 0 {
            newL := make([]int, len(p))
            copy(newL, p)
            result = append(result, newL)
            return
        }
        if i >= len(c) || t < 0 {return}
        
        // logic
        // not choose
        dfs(c, i+1, t, p)
        
        // choose
        p = append(p, c[i])
        dfs(c, i, t-c[i], p)
        p = p[:len(p)-1]
    }
    dfs(candidates, 0, target, nil)
    return result
}
