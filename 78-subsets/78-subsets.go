/*
    approach 1: 0/1 recursion
    - choose / not choose
    - coding out the choose / not choose decision tree
    - we can only use the 1 num ONCE
    - At leaves, we will append the running path
    - ( new path each recursion or backtrack existing )

*/

func subsets(nums []int) [][]int {
    var result [][]int
    var dfs func(i int, paths []int)
    dfs = func(i int, paths []int) {
        
        
        // base
        if i == len(nums) {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        
        // logic
        
        // not choose
        dfs(i+1, paths)
        
    
        // choose
        // action
        paths = append(paths, nums[i])
        // recurse
        dfs(i+1, paths)
        // backtrack
        paths = paths[:len(paths)-1]
    }
    
    dfs(0, nil)
    return result
    
}