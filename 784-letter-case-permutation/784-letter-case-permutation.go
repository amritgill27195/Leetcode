
// PUKE.... horrible horrible -- TODO: revisit this dumbass
func letterCasePermutation(s string) []string {
    seen := map[string]bool{}
    var result []string
    var helper func(i int, path string)
    helper = func(i int, path string)  {
        // base
        if i == len(s) {
            if _, ok := seen[path]; !ok {
                seen[path] = true
                result = append(result, path)
            }
            return
        }
        
        
        // logic
        strChar := string(s[i])
        upper := strings.ToUpper(strChar)
        lower := strings.ToLower(strChar)
        
        helper(i+1, path+upper)
        helper(i+1, path+lower)
    }
    helper(0, "")
    // for k , _ := range res {
    //     result = append(result, k)
    // }
    return result
}