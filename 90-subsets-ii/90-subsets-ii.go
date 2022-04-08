func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
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
            if i != start && nums[i] == nums[i-1] {continue}
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