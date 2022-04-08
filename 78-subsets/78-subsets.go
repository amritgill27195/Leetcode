/*
    approach 1: 0/1 recursion
    - choose / not choose
    - coding out the choose / not choose decision tree
    - we can only use the 1 num ONCE
    - At leaves, we will append the running path
    - ( new path each recursion or backtrack existing )
    - time: exponential ( o(2^n) where n is the number of elements in nums array )
    - space: exponential
    
    approach 2: for loop based recursion
    - instead of looking for an answer at a leaf node,
    - each path at every node is answer
*/

func subsets(nums []int) [][]int {
    
    var result [][]int
    var dfs func(start int, paths []int)
    dfs = func(start int, paths []int) {
        
        // base
        
        newL := make([]int, len(paths))
        copy(newL, paths)
        result = append(result, newL)
        if start >= len(nums) {return}
        
        
        // logic
        for i := start; i < len(nums); i++ {
            // action
            paths = append(paths, nums[i])
            // recurse
            dfs(i+1, paths)
            // backtrack
            paths = paths[:len(paths)-1]
        }
        
    }
    dfs(0, []int{})
    return result
    
}
    
// func subsets(nums []int) [][]int {
//     var result [][]int
//     var dfs func(i int, paths []int)
//     dfs = func(i int, paths []int) {
        
        
//         // base
//         if i == len(nums) {
//             newL := make([]int, len(paths))
//             copy(newL, paths)
//             result = append(result, newL)
//             return
//         }
        
//         // logic
        
//         // not choose
//         dfs(i+1, paths)
        
    
//         // choose
//         // action
//         paths = append(paths, nums[i])
//         // recurse
//         dfs(i+1, paths)
//         // backtrack
//         paths = paths[:len(paths)-1]
//     }
    
//     dfs(0, nil)
//     return result
    
// }