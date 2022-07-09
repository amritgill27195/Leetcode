func letterCombinations(digits string) []string {
    if len(digits) == 0 {return nil} 
    result := []string{}
    digMap := map[byte][]string{
        '1':[]string{},
        '2':[]string{"a","b", "c"},
        '3':[]string{"d","e", "f"},
        '4':[]string{"g","h", "i"},
        '5':[]string{"j","k", "l"},
        '6':[]string{"m","n", "o"},
        '7':[]string{"p","q", "r", "s"},
        '8':[]string{"t","u", "v"},
        '9':[]string{"w","x", "y", "z"},
    }
    var dfs func(path string, start int)
    dfs = func(path string, start int) {
        // base
        if start == len(digits) {
            result = append(result, path)
            return
        }
        
        
        // logic
        chars := digMap[digits[start]]
        for i := 0; i < len(chars); i++ {
            dfs(path+chars[i], start+1)
        }
    }
    dfs("", 0)
    return result
}