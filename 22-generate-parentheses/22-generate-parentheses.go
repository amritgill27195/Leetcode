
func generateParenthesis(n int) []string {
    result := []string{}
    var backtrack func(openN, closeN int, path string)
    backtrack = func(openN, closeN int, path string) {
        // base
        if openN == n && closeN == n {
            result = append(result, path)
            return
        }
        
        // logic
        if openN < n {
            path += "("
            backtrack(openN+1, closeN, path)
            path = path[:len(path)-1]
        }
        if closeN < openN {
            path += ")"
            backtrack(openN, closeN+1, path)
            path = path[:len(path)-1]
        }
    }
    backtrack(0,0,"")
    return result
}




// /*
//     approach: dfs and backtracking to generate all permutations of a string s
//     then checking at the end if its valid, only then adding to result if it is valid
    
//     Result: Fail, TLE
// */
// func generateParenthesis(n int) []string {
   
//     // generate initial string, it may be invalid
//     s := []string{}
//     for i := 0; i < n; i++ {
//         s = append(s, ")")
//         s = append(s, "(")
//     }
    
//     result := map[string]bool{}
//     var backtrack func(start int)
//     // then generate all permutations and check if they are balanced
//     backtrack = func(start int) {
//         // base
//         if start == len(s) {
//             joined := strings.Join(s, "")
//             if isBalanced(joined) {
//                 result[joined] = true
//             }
//             return
//         }
        
//         // logic
//         for i := start; i < len(s); i++ {
//             s[i],s[start] = s[start],s[i]
//             backtrack(start+1)
//             s[i],s[start] = s[start],s[i]
//         }
//     }
//     backtrack(0)
//     r := []string{}
//     for k, _ := range result {
//         r = append(r, k)
//     }
//     return r
// }

// func isBalanced(s string) bool {
//     count := 0
//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' {count++}
//         if s[i] == ')' {count--}
//         if count < 0 {return false}
//     }
//     return count == 0
// }


