func combine(n int, k int) [][]int {
    result := [][]int{}
    
    var backtrack func(start int, paths []int)
    backtrack = func(start int, paths []int) {
        
        // base
        if len(paths) == k {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        if start > n {return}
        
        // logic
        for i := start; i <= n; i++ {
            // action
            paths = append(paths, i)
            // recurse
            backtrack(i+1, paths)
            // backtrack
            paths = paths[:len(paths)-1]
        }
        
    }
    backtrack(1,nil)
    
    return result
}