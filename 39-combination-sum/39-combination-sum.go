// func combinationSum(candidates []int, target int) [][]int {
    
    
//     var result [][]int
//     var helper func(c []int, target int, i int, paths []int)
//     helper = func(c []int, target int, i int, paths []int){
//         // base
//         if i >= len(c) || target < 0 {return}
        
//         if target == 0 {
//             newPaths := make([]int, len(paths))
//             copy(newPaths, paths)
//             result = append(result, newPaths)
//             return
//         }
        
//         // logic
//         // not choose
//         helper(c, target, i+1, paths)
        
//         // choose
//         paths = append(paths, c[i])
//         helper(c, target - c[i], i, paths)
        
//         paths = paths[:len(paths)-1]
//     }
    
//     helper(candidates, target, 0, []int{})
//     return result
// }



func combinationSum(candidates []int, target int) [][]int {
    
    
    var result [][]int
    var helper func(c, paths []int, t, start int)
    helper = func(c, paths []int, t, start int) {
        
        // base
        if t < 0 {return}
        if t == 0 {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
    
        
        // logic
        for i := start; i < len(c); i++ {
            paths = append(paths, c[i])
            nt := t-c[i]
            helper(c, paths, nt , i)
            paths = paths[:len(paths)-1]
        }
        
    }
    helper(candidates, []int{}, target, 0)
    return result
}