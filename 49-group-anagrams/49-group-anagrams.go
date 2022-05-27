func groupAnagrams(strs []string) [][]string {
    out := [][]string{}
    
    m := map[string][]string{}
    for i := 0; i < len(strs); i++ {
        str := strs[i]
        strSplit := strings.Split(str,"")
        sort.Strings(strSplit)
        strJoined := strings.Join(strSplit,"")
        m[strJoined] = append(m[strJoined], str)
    }
    
    for _, v := range m {
        out = append(out, v)
    }
    
    return out
}