func numTilePossibilities(tiles string) int {
    if len(tiles) <= 1 {return len(tiles)}
    
    freqMap := map[byte]int{}
    for i := 0; i < len(tiles); i++ { freqMap[tiles[i]]++ }
    
    count := 0
    var dfs func(path string) 
    dfs = func(path string) {
        // base
        if path != "" {
            count++
        }
        
        // logic
        for k, v := range freqMap {
            if v == 0 {continue}
            // action
            freqMap[k]--
            path += string(k)
            // recurse
            dfs(path)
            // backtrack
            freqMap[k]++
        }
    }
    dfs("")
    return count
    
}