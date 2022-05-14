
type AutocompleteSystem struct {
	freqMap map[string]int
	input   *strings.Builder
	result  []string
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	m := map[string]int{}
	for i := 0; i < len(sentences); i++ {
		m[sentences[i]] += times[i]
	}
	return AutocompleteSystem{freqMap: m, input: new(strings.Builder), result: []string{}}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.freqMap[this.input.String()]++
		this.input = new(strings.Builder)
		return nil
	}
	this.input.WriteByte(c)
	for k, _ := range this.freqMap {
		in := this.input.String()
		if strings.HasPrefix(k, in) {
			heap.Push(this, k)
			if len(this.result) > 3 {
				heap.Pop(this)
			}
		}
	}
	out := []string{}
	for len(this.result) != 0 {
		out = append(out, heap.Pop(this).(string))
	}
	for i := 0; i < len(out)/2; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

// Pop : remove and return element Len() - 1.a
func (acs *AutocompleteSystem) Pop() interface{} {
	out := acs.result[len(acs.result)-1]
	acs.result = acs.result[:len(acs.result)-1]
	return out
}
func (acs *AutocompleteSystem) Swap(i, j int) {
	acs.result[i], acs.result[j] = acs.result[j], acs.result[i]
}
func (acs *AutocompleteSystem) Len() int           { return len(acs.result) }
func (acs *AutocompleteSystem) Push(x interface{}) { acs.result = append(acs.result, x.(string)) }
func (acs *AutocompleteSystem) Less(i, j int) bool {
	iWord := acs.result[i]
	iWordWeight := acs.freqMap[iWord]
	jWord := acs.result[j]
	jWordWeight := acs.freqMap[jWord]

	if iWordWeight == jWordWeight {
		return strings.Compare(jWord, iWord) == -1
	}
	return iWordWeight < jWordWeight

}
