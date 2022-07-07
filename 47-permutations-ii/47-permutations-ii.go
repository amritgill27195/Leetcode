func permuteUnique(nums []int) [][]int {
    result := [][]int{}
    cache := map[string]bool{}
    var dfs func(start int)
    dfs = func(start int) {
        // logic
        seenData, _ := json.Marshal(nums)
        if cache[string(seenData)] { return }
        
        if start == len(nums) {
            newL := make([]int, len(nums))
            copy(newL, nums)
            result = append(result, newL)
            return
        }
        
        // base
        for i := start; i < len(nums); i++ {
            // action
            nums[start],nums[i] = nums[i], nums[start]
            // recurse
            dfs(start+1)
            // backtrack
            nums[start],nums[i] = nums[i], nums[start]
        }
        
        data, _ := json.Marshal(nums)
        cache[string(data)] = true
    }
    
    dfs(0)
    return result
}