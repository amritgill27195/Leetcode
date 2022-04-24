// worst problem desc ever...
// brute force
// doing this question brute force first makes me understand why binary search helps...
// time: o(c) where c is the number of citations
// func hIndex(citations []int) int { 
//     n := len(citations)
//     for i := 0; i < len(citations); i++ {
//         diff := n-i
//         if diff <= citations[i] {
//             return diff
//         }
//     }
//     return 0
// }


// time: o(logn)
// space: o(1)
func hIndex(citations []int) int {
    left := 0
    hIdx := 0
    n := len(citations)
    right := len(citations)-1
    for left <= right {
        mid := left + (right-left)/2
        if n-mid <= citations[mid] {
            hIdx = n-mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return hIdx
}