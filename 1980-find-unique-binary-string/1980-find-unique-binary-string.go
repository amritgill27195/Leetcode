func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    
    set := map[string]bool{}
    for i := 0; i < n; i++ {
        set[nums[i]] = true   
    }
    m := map[string]int{"1":n, "0":n}
    
    var dfs func(path string) string
    dfs = func(path string) string {
        // base
        if len(path) == n {
            if _, ok := set[path]; !ok {
                return path
            }
            return ""
        }
        
        // logic
        for k, v := range m {
            if v == 0 {continue}
            path += k
            m[k]--
            if found := dfs(path); found != "" {
                return found
            }
            path = path[:len(path)-1]
            m[k]++
        }
        return ""
    }
    return dfs("") 
}