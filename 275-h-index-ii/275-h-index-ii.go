// worst problem desc ever...
// time: o(logn)
// space: o(1)
// func hIndex(citations []int) int {
//     left := 0
//     right := len(citations)-1
//     n := len(citations)
//     for left <= right {
//         mid := left+(right-left)/2
//         if n-mid == citations[mid] {
//             return n-mid
//         } else if citations[mid] > n-mid {
//             right = mid-1
//         } else {
//             left = mid+1
//         }
//     }
//     return n-left
// }


func hIndex(citations []int) int { 
    n := len(citations)
    for i := 0; i < len(citations); i++ {
        diff := n-i
        if diff <= citations[i] {
            return diff
        }
    }
    return 0
}

func abs(n int) int {
    if n < 0 {return n * -1}
    return n
}