
func letterCasePermutation(s string) []string {
    var result []string
    var backtrack func(i int, path string)
    backtrack = func(i int, path string)  {
        // base
        if i == len(s) {
            result = append(result, path)
            return
        }
        
        
        // logic
        char := s[i]
        strChar := string(char)
        if char >= '0' && char <= '9' {
            backtrack(i+1, path+strChar)
        } else {
            backtrack(i+1, path+strings.ToUpper(strChar))
            backtrack(i+1, path+strings.ToLower(strChar))
        }
    }
    backtrack(0, "")
    return result
}