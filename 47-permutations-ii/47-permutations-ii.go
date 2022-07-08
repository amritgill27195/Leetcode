func permuteUnique(nums []int) [][]int {
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
    }
    result := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newL := make([]int, len(path))
            copy(newL, path)
            result = append(result, newL)
            return
        }
        
        // logic
        for k, v := range freqMap {
            if v == 0 {continue}
            // action
            path = append(path, k)
            freqMap[k]--
            // recurse
            dfs(path)
            // backtrack
            freqMap[k]++
            path = path[:len(path)-1]
        }
    }
    dfs(nil)
    return result
}