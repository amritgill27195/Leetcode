func findClosestElements(arr []int, k int, x int) []int {
	if arr == nil {
		return nil
	}
	mx := &maxHeap{items: []*item{}}
	for _, e := range arr {
		i := &item{val: e, distAway: int(math.Abs(float64(e-x)))}
		heap.Push(mx, i)
		if len(mx.items) > k {
			heap.Pop(mx)
		}
	}
	out := make([]int, k)
	for i := len(out)-1 ; i >= 0; i-- {
		out[i]= heap.Pop(mx).(*item).val
	}
	sort.Ints(out)
	return out
}

type item struct {
	distAway int
	val int
}

type maxHeap struct {
	items []*item
}
func (mn *maxHeap) Less(i, j int) bool {
	if mn.items[i].distAway == mn.items[j].distAway {
		return mn.items[i].val > mn.items[j].val
	}
	return mn.items[i].distAway > mn.items[j].distAway
}
func (mn *maxHeap) Swap(i, j int) { mn.items[i] , mn.items[j] = mn.items[j] , mn.items[i]}
func (mn *maxHeap) Len() int { return len(mn.items) }
func (mn *maxHeap) Push(x interface{}) { mn.items = append(mn.items, x.(*item))}
func (mn *maxHeap) Pop() interface{} {
	out := mn.items[len(mn.items)-1]
	mn.items = mn.items[:len(mn.items)-1]
	return out
}