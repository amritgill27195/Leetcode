import (
    "sort"
    "strings"
)
// map solution
// time: o(s+t)
// space: o(min(s,t))
// func isAnagram(s string, t string) bool {
//     // make sure s is the smaller string
//     if len(t) < len(s) {
//         return isAnagram(t,s)
//     }
    
//     sMap := map[string]int{}
//     for _, char := range s{
//         sMap[string(char)]++
//     }
    
//     for _, char := range t {
//         tChar := string(char)
//         val, ok := sMap[tChar]
//         if !ok {
//             return false
//         } else if val == 0 {
//             return false
//         }
//         sMap[tChar]--
//     }
//     return true
// }



// sort both strings and return equality check
func isAnagram(s string, t string) bool {
    news := strings.Split(s, "")
    sort.Strings(news)
    s = strings.Join(news, "")
    newt := strings.Split(t, "")
    sort.Strings(newt)
    t = strings.Join(newt, "")
    return s == t
}