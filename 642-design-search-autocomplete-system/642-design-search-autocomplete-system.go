type AutocompleteSystem struct {
    freqMap map[string]int
    input *strings.Builder
}


func Constructor(sentences []string, times []int) AutocompleteSystem {
    m := map[string]int{}
    for i := 0; i < len(sentences); i++ {
        m[sentences[i]] += times[i] 
    }
    return AutocompleteSystem{freqMap: m,input: new(strings.Builder)} 
}




func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        this.freqMap[this.input.String()]++
        this.input = new(strings.Builder)
        return nil
    }
    this.input.WriteByte(c)
    type ansPair struct{
        word string
        times int
    }
    ans := []*ansPair{}
    for k, v := range this.freqMap {
        in := this.input.String()
        if strings.HasPrefix(k, in) {
            ans = append(ans, &ansPair{word: k, times: v})
        }
    }
    
    sort.Slice(ans, func(i, j int)bool {
        if ans[i].times == ans[j].times {
            return ans[i].word < ans[j].word
        }
        return ans[i].times > ans[j].times
    })
   
    out := []string{}
    for i := 0; i < len(ans); i++ {
        out = append(out, ans[i].word)
        if len(out) == 3 {
            break
        }
    }
    return out
}


/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */