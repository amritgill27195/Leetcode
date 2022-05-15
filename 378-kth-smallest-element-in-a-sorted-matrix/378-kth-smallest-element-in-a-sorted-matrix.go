func kthSmallest(matrix [][]int, k int) int {
    mx := &maxHeap{items: []int{}}
    m := len(matrix)
    n := len(matrix[0])
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            heap.Push(mx, matrix[i][j])
            if len(mx.items) > k {
                heap.Pop(mx)
            }
        }
    }
    return mx.items[0]
}



type maxHeap struct {
	items []int
}


func (mx *maxHeap) Len() int {return len(mx.items)}
func (mx *maxHeap) Swap(i, j int) {mx.items[i],mx.items[j]= mx.items[j],mx.items[i]}
func (mx *maxHeap) Less(i, j int) bool {return mx.items[i] > mx.items[j]}
func (mx *maxHeap) Pop() interface{} {
	out := mx.items[len(mx.items)-1]
	mx.items = mx.items[:len(mx.items)-1]
	return out
}
func (mx *maxHeap) Push(x interface{}) {mx.items = append(mx.items, x.(int))}
