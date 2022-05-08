
/*
    approach: brute force
    - using a freqMap
*/
type AutocompleteSystem struct {
    input *strings.Builder
    m map[string]int
}


func Constructor(sentences []string, times []int) AutocompleteSystem {
    m := map[string]int{}
    for i := 0; i < len(sentences); i++ {
        sentence := sentences[i]
        freq := times[i]
        m[sentence] += freq
    }
    return AutocompleteSystem{
        m: m,
        input: new(strings.Builder),
    }
}


// time: o(nk) + o(nlogn)
// avg time: o(nlogn)
// space: o(n) 
func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        this.m[this.input.String()]++
        this.input = new(strings.Builder)
        return nil
    }
    this.input.WriteByte(c)
    
    type pair struct {
        s string
        t int
    }
    
    tmp := []*pair{}
    // time: o(n) -- we loop thru entire map AND
    for key, val := range this.m {
        // time: o(k) - where k is the length of this.input...
        if strings.HasPrefix(key, this.input.String()) {
            tmp = append(tmp, &pair{s: key, t: val})
        }
    }
    // so above time is o(nk) - k will likely be a smaller term on avg so o(n) on avg
    
    // sorting the entire tmp result -- o(nlogn)
    sort.Slice(tmp, func(i,j int)bool{
        if tmp[i].t == tmp[j].t {
            return tmp[i].s < tmp[j].s
        }
        return tmp[i].t > tmp[j].t
    })
    
    // time: o(1) -- its constant -- 3 -- we break as soon as we have 3 results...
    result := []string{}
    for j := 0; j < len(tmp); j++ {
        result = append(result, tmp[j].s)
        if len(result) == 3 {
            break
        }
    }
    return result
}


/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */