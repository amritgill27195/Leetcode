/*
    approach: store a pair(val, distAway) in a maxHeap
    - For each n elements
    - Calc the distance : arr[n] - x
    - Create a pair {val, distAway}
    - Store in maxHeap ( maxHeap sorted by distAway so the farthest pair will be on the top )
    - If the maxHeap size > k heap.pop()
    - Once all the elements are done,
    - The maxHeap will be left with k pairs that are the closest to x
    - Loop over maxHeap items and store them in result array and return
    
    time: o(n) * o(logk) = o(nlogk)
    - why k? because in maxHeap we will only store k elements
    - where did the n come from? We loop over all elements to store them in maxHeap therefore o(n) * o(logk)
    space: o(k) - in maxHeap at worse, we will store k elements
*/


// maxHeap impl
type pair struct {
    dist int
    val int
}
func findClosestElements(arr []int, k int, x int) []int {
    if arr == nil || len(arr) == 0 {return nil}
    
    mx := &maxHeap{items: []*pair{}}
    for i := 0; i < len(arr); i++ {
        p := &pair{dist: abs(x-arr[i]), val: arr[i]}
        heap.Push(mx, p)
        if len(mx.items) > k {
            heap.Pop(mx)
        }
    }
    out := make([]int, len(mx.items))
    for i := len(out)-1; i >= 0; i-- {
        out[i] = heap.Pop(mx).(*pair).val
    }
    sort.Ints(out)
    return out
    
}
func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}


type maxHeap struct {
	items []*pair
}

func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Less(i, j int) bool {
    if m.items[i].dist == m.items[j].dist {
        return m.items[i].val > m.items[j].val
    }
    return m.items[i].dist > m.items[j].dist
    
}
func (m *maxHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(*pair))}
func (m *maxHeap) Pop() interface{} {
	i := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return i
}


