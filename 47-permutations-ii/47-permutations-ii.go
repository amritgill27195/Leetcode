func permuteUnique(nums []int) [][]int {
    
    result := [][]int{}
    m := map[int]int{}
    for _, n := range nums {
        m[n]++
    }
    
    var dfs func(paths []int)
	dfs = func(paths []int) {
            
        // base
        if len(paths) == len(nums) {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
		}
        
        // logic
        for num, freq := range m {
            if freq > 0 {
                paths = append(paths, num)
                m[num]--
                dfs(paths)
                paths = paths[:len(paths)-1]
                m[num]++
            }
		}
	}
	dfs(nil)
	return result    
}