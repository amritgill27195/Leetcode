func hIndex(citations []int) int {
    sort.Ints(citations)
    left := 0
    right := len(citations)-1
    n := len(citations)
    hIdx := 0
    for left <= right {
        mid := left + (right-left)/2
        diff := n-mid
        if diff <= citations[mid] {
            if diff == citations[mid] {
                return diff
            }
            hIdx = diff
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return hIdx
}