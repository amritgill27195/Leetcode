func letterCasePermutation(s string) []string {
    var result []string
    var dfs func(start int, path string)
    dfs = func(start int, path string) {
        // fmt.Println("Path: ", path)
        // base
        if start == len(s) {
            if len(path) == len(s) {
                result = append(result, path)
            }
            return
        }
        
        //logic
        for i := start; i < len(s); i++ {
            char := s[i]
            strChar := string(char)
            if char >= '0' && char <= '9' {
                dfs(i+1, path+strChar)
            } else {
                dfs(i+1, path+strings.ToLower(strChar))
                dfs(i+1, path+strings.ToUpper(strChar))
            }
        }
    }
    dfs(0, "")
    return result
}