func permute(nums []int) [][]int {
    var result [][]int
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) {
            newL := make([]int, len(nums))
            copy(newL, nums)
            result = append(result, newL)
            return
        }
        // logic
        for i := start; i < len(nums); i++ {
            // action
            nums[i], nums[start] = nums[start], nums[i]
            // recurse
            dfs(start+1)
            // backtrack
            nums[i], nums[start] = nums[start], nums[i]
        }
    }
    
    dfs(0)
    return result
}