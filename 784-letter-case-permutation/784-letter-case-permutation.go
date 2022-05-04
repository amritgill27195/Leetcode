
// PUKE.... horrible horrible -- TODO: revisit this dumbass
func letterCasePermutation(s string) []string {
    res := map[string]bool{}
    var result []string
    var helper func(i int, path string)
    helper = func(i int, path string)  {
        // base
        if len(path) == len(s) {
            res[path] = true
            return
        }
        if i == len(s) {return}
        
        
        // logic
        upper := strings.ToUpper(string(s[i]))
        lower := strings.ToLower(string(s[i]))
        helper(i+1, path+upper)
        helper(i+1, path+lower)
    }
    helper(0, "")
    for k , _ := range res {
        result = append(result, k)
    }
    return result
}