func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    var backtrack func(start, rs int, path []int)
    backtrack = func(start, rs int, path []int) {
        // base
        if rs == target {
            newL := make([]int, len(path))
            copy(newL, path)
            result = append(result, newL)
            return
        }
        if rs > target {return}
        
        // logic
        for i := start; i < len(candidates); i++ {
            // action
            path = append(path, candidates[i])
            // recurse
            backtrack(i, rs+candidates[i], path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    backtrack(0, 0, nil)
    return result
}