func letterCombinations(digits string) []string {
    if len(digits) == 0 {return nil}
    charsMap := map[byte][]string{
        '2' : []string{"a","b","c"},
        '3' : []string{"d","e","f"},
        '4' : []string{"g","h","i"},
        '5' : []string{"j","k","l"},
        '6' : []string{"m","n","o"},
        '7' : []string{"p","q","r", "s"},
        '8' : []string{"t","u","v"},
        '9' : []string{"w","x","y", "z"},
    }
    result := []string{}
    var backtrack func(start int, path string)
    backtrack = func(start int, path string) {
        // base
        if len(path) == len(digits) {
            result = append(result, path)
            return
        }
        if start == len(digits) {return}
        
        // logic
        chars := charsMap[digits[start]]
        for j := 0; j < len(chars); j++ {
            // action
            path += chars[j]
            // recurse
            backtrack(start+1, path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    backtrack(0, "")
    return result
}