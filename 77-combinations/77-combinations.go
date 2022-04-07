func combine(n int, k int) [][]int {
    result := [][]int{}
    
    var backtrack func(start int,end int, paths []int)
    backtrack = func(start int, end int, paths []int) {
        
        // base
        if len(paths) == k {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        if start > n {return}
        
        // logic
        for i := start; i <= end; i++ {
            // action
            paths = append(paths, i)

            // recurse
            backtrack(i+1, end, paths)
            // backtrack
            paths = paths[:len(paths)-1]
        }
        
    }
    backtrack(1,n,nil)
    
    return result
}