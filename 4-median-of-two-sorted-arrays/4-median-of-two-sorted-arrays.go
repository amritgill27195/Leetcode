
/*
    approach : merge + sort and then grab middle
    - merge both arrays into a new m+n array
    - Sort the new array
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    time: o(m+nlogm+n)
    space: o(1)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     for i := 0; i < len(nums1); i++ {
//         merged = append(merged, nums1[i])
//     }
//     for i := 0; i < len(nums2); i++ {
//         merged = append(merged, nums2[i])
//     }
//     sort.Ints(merged)
//     midIdx := len(merged) / 2
//     if len(merged) % 2 != 0 {
//         // odd array len, return the mid
//         return float64(merged[midIdx])
//     }
//     var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
//     return avg
// }


/*
    approach : merge using two pointers so no need to sort, and then grab middle
    - merge both arrays into a new m+n array using 2 pointers
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    time: o(m+nlogm+n)
    space: o(1)
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] < nums2[n2] {
            merged = append(merged, nums1[n1])
            n1++
        } else {
            merged = append(merged, nums2[n2])
            n2++
        }
    }
    if n1 < len(nums1) {
        merged = append(merged, nums1[n1:]...)
    }
    if n2 < len(nums2) {
        merged = append(merged, nums2[n2:]...)
    }
    fmt.Println(merged)
    
    midIdx := len(merged) / 2
    if len(merged) % 2 != 0 {
        // odd array len, return the mid
        return float64(merged[midIdx])
    }
    var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
    return avg
}