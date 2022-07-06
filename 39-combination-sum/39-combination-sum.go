func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    var dfs func(start int, path []int, sum int)
    dfs = func(start int, path []int, sum int) {
        // base
        if start == len(candidates) || sum > target {
            return
        }
        if sum == target {
            newL := make([]int, len(path))
            copy(newL, path)
            result = append(result, newL)
            return
        }
        
        // logic
        for i := start; i < len(candidates); i++ {
            // action
            path = append(path, candidates[i])
            // recurse
            dfs(i, path, sum+candidates[i])
            // backtrack
            path = path[:len(path)-1]
        }
    }
    dfs(0, nil, 0)
    return result
}