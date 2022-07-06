func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    var dfs func (start int, t int, path []int)
    dfs = func(start int, t int, path []int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                result = append(result, newL)
            }
            return
        }
        
        // logic
        for i := start; i < len(candidates); i++ {
            // action 
            path = append(path, candidates[i])
            // recurse
            dfs(i,t-candidates[i],path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    dfs(0, target, nil)
    return result
}