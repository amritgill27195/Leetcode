func canJump(nums []int) bool {
    memo := map[int]bool{}
    var dfs func(idx int) bool
    dfs = func(idx int) bool {
        // base
        if idx >= len(nums)-1 {return true}
        
        // logic
        for i := nums[idx]; i >= 1; i-- {
            nextIdx := idx + i
            if val, ok := memo[nextIdx]; ok {
                if val == true{return true}
            } else {
                res := dfs(nextIdx)
                memo[nextIdx] = res
                if res {return true}
            }
        }
        return false
    }
    return dfs(0)
}