func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    
    for i := 0; i < len(strs); i++ {
        word := strs[i]
        wordSplit := strings.Split(word, "")
        sort.Strings(wordSplit)
        wordSorted := strings.Join(wordSplit, "")
        m[wordSorted] = append(m[wordSorted],word)
    }
    out := [][]string{}
    for _, v := range m {
        out = append(out, v)
    }
    return out
}