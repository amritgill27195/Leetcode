// using same technique as for loop based recursion
// and using hashmap + index based deduped array to perform for loop based recursion on
func subsetsWithDup(nums []int) [][]int {
    if nums == nil || len(nums) == 0 {
        return nil
    }
    
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ { freqMap[nums[i]]++ }
    
    deduped := [][]int{}
    for k, v := range freqMap { deduped = append(deduped, []int{k, v})}
    
    result := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        newL := make([]int, len(path))
        copy(newL, path)
        result = append(result, newL)
        
        // logic
        for i := start;i<len(deduped);i++ {
            if deduped[i][1] == 0 {continue}
            // action
            path = append(path, deduped[i][0])
            deduped[i][1]--
            // recurse
            dfs(i, path)
            // backtrack
            path = path[:len(path)-1]
            deduped[i][1]++
        }
    }
    dfs(0, nil)
    return result
}