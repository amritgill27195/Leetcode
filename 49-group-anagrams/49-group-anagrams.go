func groupAnagrams(strs []string) [][]string {
    m := map[float64][]string{}
    for i := 0 ; i < len(strs); i++ {
        word := strs[i]
        hashVal := hash(word)
        m[hashVal] = append(m[hashVal], strs[i])
    }
    out := [][]string{}
    for _, v := range m {
        out = append(out, v)
    }
    
    return out
   
}

var primes = []float64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}

func hash(word string) float64 {
    prod := 1.0
    for i := 0; i < len(word); i++ {
        asciValOfChar := word[i] // e -- byte
        charIdx := asciValOfChar - 'a'
        prod *= primes[charIdx]
    }
    return prod
}