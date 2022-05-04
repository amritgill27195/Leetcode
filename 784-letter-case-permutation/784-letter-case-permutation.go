
func letterCasePermutation(s string) []string {
    var result []string
    var helper func(i int, path string)
    helper = func(i int, path string)  {
        // base
        if i == len(s) {
            result = append(result, path)
            return
        }
        
        
        // logic
        char := s[i]
        strChar := string(char)
        if char >= '0' && char <= '9' {
            helper(i+1, path+strChar)
        } else {
            helper(i+1, path+strings.ToUpper(strChar))
            helper(i+1, path+strings.ToLower(strChar))
        }
    }
    helper(0, "")
    return result
}