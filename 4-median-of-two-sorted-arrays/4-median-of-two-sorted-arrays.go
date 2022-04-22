
/*
    approach : merge + sort and then grab middle
    - merge both arrays into a new m+n array
    - Sort the new array
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    
    
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    for i := 0; i < len(nums1); i++ {
        merged = append(merged, nums1[i])
    }
    for i := 0; i < len(nums2); i++ {
        merged = append(merged, nums2[i])
    }
    sort.Ints(merged)
    midIdx := len(merged) / 2
    if len(merged) % 2 != 0 {
        // odd array len, return the mid
        return float64(merged[midIdx])
    }
    var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
    return avg
}