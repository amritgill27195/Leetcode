func combinationSum2(candidates []int, target int) [][]int {
    result := [][]int{}
    
    freqMap := map[int]int{}
    for i := 0; i < len(candidates); i++ {
        freqMap[candidates[i]]++
    }
    
    deduped := [][]int{}
    for k, v := range freqMap {
        deduped = append(deduped, []int{k,v})
    }
    
    var dfs func(start int, path []int, t int)
    dfs = func(start int, path []int, t int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                result = append(result, newL)
            }
            return
        }
        
        // logic
        for i := start; i < len(deduped); i++ {
            val := deduped[i][0]
            count := deduped[i][1]
            if count == 0 {continue}
            
            // action
            path = append(path, val)
            deduped[i][1]--
            // recurse
            dfs(i,path,t-val)
            // backtrack
            path = path[:len(path)-1]
            deduped[i][1] = count
        
        }
    }
    
    dfs(0,nil, target)
    return result
    
}