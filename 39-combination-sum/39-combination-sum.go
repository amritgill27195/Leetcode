func combinationSum(candidates []int, target int) [][]int {
    
    
    var result [][]int
    var helper func(c []int, target int, i int, paths []int)
    helper = func(c []int, target int, i int, paths []int){
        // base
        if i >= len(c) || target < 0 {return}
        
        if target == 0 {
            newPaths := make([]int, len(paths))
            copy(newPaths, paths)
            result = append(result, newPaths)
            return
        }
        
        // logic
        // not choose
        helper(c, target, i+1, paths)
        
        // choose
        paths = append(paths, c[i])
        helper(c, target - c[i], i, paths)
        
        paths = paths[:len(paths)-1]
    }
    
    helper(candidates, target, 0, []int{})
    return result
}