func groupAnagrams(strs []string) [][]string {
    
    out := [][]string{}
    
    m := map[float64][]string{}
    for i := 0; i < len(strs); i++ {
        str := strs[i]
        prod := hash(str)
        m[prod] = append(m[prod], str)
    }
    
    for _, v := range m {
        out = append(out, v)
    }
    
    return out
}


var primes = []float64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}
func hash(s string) float64 {
    primeProd := 1.0
    for _, char := range s {
        primeProd *= primes[char-'a']
    }
    return primeProd
}