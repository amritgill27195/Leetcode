func letterCombinations(digits string) []string {
    if len(digits) == 0 {return nil}
    digMap := map[byte][]string{
        '2': []string{"a","b","c"},
        '3': []string{"d", "e", "f"},
        '4': []string{"g", "h", "i"},
        '5': []string{"j", "k", "l"},
        '6': []string{"m","n","o"},
        '7': []string{"p", "q", "r","s"},
        '8': []string{"t","u","v"},
        '9': []string{"w","x","y", "z"},
    }
    var result []string
    var dfs func(path string, start int)
    dfs = func(path string, start int) {
        // base
        if len(path) == len(digits) {
            result = append(result, path)
            return
        }
        
        
        // logic
        for i := start; i < len(digits); i++ {
            for _, char := range digMap[digits[i]] {
                path += string(char)
                dfs(path, i+1)
                path = string(path[:len(path)-1])
            }
        }
    }
    dfs("",0)
    return result
}