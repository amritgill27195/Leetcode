func findSubsequences(nums []int) [][]int {
    var result [][]int
    var backtrack func(start int, paths []int)
    backtrack = func(start int, paths []int) {
        // base
        if len(paths) >= 2 {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
        }
        
        // logic
        vis := map[int]bool{}
        for i := start; i < len(nums); i++ {
            if (len(paths) == 0 || nums[i] >= paths[len(paths)-1]) && (!vis[nums[i]]){
                // action
                vis[nums[i]] = true
                paths = append(paths, nums[i])
                // recurse
                backtrack(i+1, paths)
                // backtrack
                paths = paths[:len(paths)-1]
            }
        }
    }
    backtrack(0, nil)
    return result
}