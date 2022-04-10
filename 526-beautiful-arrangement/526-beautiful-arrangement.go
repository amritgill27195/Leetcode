func countArrangement(n int) int {
    count := 0
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i+1
    }
    
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) {
            count++
            return
        }
        
        //logic
        for i := start; i < len(nums); i++ {
            // check
            if nums[i] % (start+1) == 0 || (start+1) % nums[i] == 0 {
                // action
                nums[i], nums[start] = nums[start], nums[i]
                // recurse
                dfs(start+1)
                // backtrack
                nums[i], nums[start] = nums[start], nums[i]
            }
        }
    }
    
    dfs(0)
    
    return count
}