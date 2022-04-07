// func combinationSum(candidates []int, target int) [][]int {
//     result := [][]int{}
    
//     var dfs func (c []int, i int, t int, p []int)
//     dfs = func(c []int,i int, t int, p []int) {
//         // base
//         if t == 0 {
//             newL := make([]int, len(p))
//             copy(newL, p)
//             result = append(result, newL)
//             return
//         }
//         if i >= len(c) || t < 0 {return}
        
//         // logic
//         // not choose
//         dfs(c, i+1, t, p)
        
//         // choose
//         p = append(p, c[i])
//         dfs(c, i, t-c[i], p)
//         p = p[:len(p)-1]
//     }
//     dfs(candidates, 0, target, nil)
//     return result
// }


func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    var backtrack func(c []int, t int, start int, paths []int)
    backtrack = func(c []int, t int, start int, paths []int) {
        
        // base 
        if t < 0 || start >= len(c) {
            return
        }
        
        if t == 0 {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        
        // logic
        for i := start; i<len(c); i++ {
            // action
            paths = append(paths, c[i])
            // recurse
            backtrack(c, t - c[i], i, paths)
            // backtrack
            paths = paths[:len(paths)-1]
        }
        
    }
    
    backtrack(candidates, target, 0, nil)
    
    return result
}
