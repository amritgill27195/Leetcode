type AutocompleteSystem struct {
    input string
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
        input: "",
    }
}


func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        this.m[this.input]++
        this.input = ""
        return nil
    }
    this.input += string(c)
    type pair struct {
        s string
        t int
    }
    tmp := []*pair{}
    for key, val := range this.m {
        if strings.HasPrefix(key, this.input) {
            tmp = append(tmp, &pair{s: key, t: val})
        }
    }
    
    sort.Slice(tmp, func(i,j int)bool{
        if tmp[i].t == tmp[j].t {
            return tmp[i].s < tmp[j].s
        }
        return tmp[i].t > tmp[j].t
    })
    
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