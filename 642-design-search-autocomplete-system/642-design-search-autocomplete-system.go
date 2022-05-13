/*
    approach: brute force
    - We need a way to make sentences and times array query-able
    - as in how many times a word repeats 
    - so to correlate a word with numberOfTimes searched, we will use a map
    - in our constructor, we will turn the input args into a freqMap
    - We also need a way to keep on appending the input string to input char
    - so we will have a strings.builder() thats appending characters of input until we the input char is #
    - when the input char is #, then we will take the string from stringBuilder and add it back to freqMap
    - Now to searching...
    - we will loop thru our freqMap and check if each word starts with our input stringBuilder
    - if it does we will add it to a list as a pair [word, numTimes]
    - Then we will sort the list by numTimes ( descending order - large number first )
        - if the numTimes for i,j are the same, then we will sort by alphabetically
        - whichever alphabet comes first will be first
    
    time: o(nk) -- to populate ans list --  + o(nlogn) -- to sort ans list --- + o(1) to populate a 3 len output list
    space: o(n) 
          
*/

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