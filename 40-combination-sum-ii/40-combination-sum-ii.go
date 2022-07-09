func combinationSum2(candidates []int, target int) [][]int {
    if candidates == nil || len(candidates) == 0 {
        return nil
    }
    freqMap := map[int]int{}
    for i := 0; i < len(candidates); i++ { freqMap[candidates[i]]++ }

    deduped := [][]int{}
    for k, v := range freqMap { deduped = append(deduped, []int{k,v}) }
    
    result := [][]int{}
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
            if deduped[i][1] == 0 {continue}
            // action
            path = append(path, deduped[i][0])
            deduped[i][1]--
            // recurse
            dfs(i, path, t-deduped[i][0])
            // backtrack
            path = path[:len(path)-1]
            deduped[i][1]++
        }
    }
    dfs(0, nil, target)
    return result
}