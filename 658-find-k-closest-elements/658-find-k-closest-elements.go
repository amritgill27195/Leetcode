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
    However I still had to sort the final result -- Did I do something wrong?  (+ o(klogk) to time for sorting the final k size output list)
    
    
    approach: 2 pointers
    - left ptr at 0
    - right ptr at n-1
    - while the distance between left and right != k
        - i.e until we do not have a window of size k
        - get leftDist to x
        - get rightDist to x
        - whoever is farther from x, move away from that value 
            - if rightDist is farther from x compared to leftDist, then right--
            - else left++
        - If the distance is the same, then move away from bigger value
    - once we have a valid window of size k
    - then loop from left to right and populate the result arr
    - time: o(n-k)+o(k)
    - space: o(1)
    
    
    approach: Binary search + two pointers
    - Since the array is sorted
    - There is a potential for binary search and/or two pointers
    - We can search for the 'x' using binary search or closest to 'x'
    - This will be in o(logn) time with o(1) space
    - Then once we have the closest to 'x', then
    - Make sure you pick closest between left vs right after binary search
    - We can start forming a window of size k
    - i.e start 2 pointers from the idx that is closest to x
    - if the dist between left and right ptr (right-left+1) == k - break - this means we have a valid window of size k 
    - if we do not, then get the 
*/


// binary search + 2 pinters
func findClosestElements(arr []int, k int, x int) []int {
    // first find the closest to x element using binary search ( o(logN) )
    left := 0
    right := len(arr)-1
    for left < right {
        mid := left + (right-left)/2
         if arr[mid] >= x {
            right = mid
        } else {
            left = mid+1
        }
    }
    left = left-1
    right = left+1

    for right-left-1 < k {
        if left == -1 {
            right++
            continue
        }
        if right == len(arr) || abs(arr[left]-x) <= abs(arr[right]-x) {
            left--
        } else {
            right++
        }
        
    }
    out := []int{}
    for i := left+1; i < right; i++ {
        out = append(out, arr[i])
    }
    return out
}


// two pointers 
// func findClosestElements(arr []int, k int, x int) []int {
//     if arr == nil || len(arr) == 0 || k == 0 {return nil}
//     left := 0
//     right := len(arr)-1
//     for right-left+1 != k {
//         leftDist := x-arr[left]
//         rightDist := arr[right]-x
//         if leftDist > rightDist{
//             left++
//         } else {
//             // if rightDist is farther away from x compared to leftDist we move right--
//             // if the distances are equal, we move away from bigger value, and in sorted array, the bigger value will ALWAYS BE RIGHT VALUE COMPARED TO LEFT VALUE
//             // therefore no need of explicit "if checks"
//             right--
//         }
//     }
    
//     out := []int{}
//     for i := left; i <= right ; i++ {
//         out = append(out, arr[i])
//     }
//     return out
// }

// maxHeap impl
// type pair struct {
//     dist int
//     val int
// }
// func findClosestElements(arr []int, k int, x int) []int {
//     if arr == nil || len(arr) == 0 {return nil}
    
//     mx := &maxHeap{items: []*pair{}}
//     for i := 0; i < len(arr); i++ {
//         p := &pair{dist: abs(x-arr[i]), val: arr[i]}
//         heap.Push(mx, p)
//         if len(mx.items) > k {
//             heap.Pop(mx)
//         }
//     }
//     out := make([]int, len(mx.items))
//     for i := len(out)-1; i >= 0; i-- {
//         out[i] = heap.Pop(mx).(*pair).val
//     }
//     // With heap I still had to sort at the end -- did I do something wrong here?
//     sort.Ints(out)
//     return out
    
// }



// type maxHeap struct {
// 	items []*pair
// }

// func (m *maxHeap) Len() int {return len(m.items)}
// func (m *maxHeap) Less(i, j int) bool {
//     if m.items[i].dist == m.items[j].dist {
//         // if distances are the same, then the bigger val takes precedence
//         return m.items[i].val > m.items[j].val
//     }
//     // or else sort 
//     return m.items[i].dist > m.items[j].dist
    
// }
// func (m *maxHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
// func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(*pair))}
// func (m *maxHeap) Pop() interface{} {
// 	i := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return i
// }



func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}