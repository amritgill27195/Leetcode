func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {
        return intersect(nums2, nums1)
    }
    
    sort.Ints(nums1)
    sort.Ints(nums2)
    
    n1 := 0
    n2 := 0
    out := []int{}
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] == nums2[n2] {
            out = append(out, nums1[n1])
            n1++
            n2++
        } else if nums1[n1] < nums2[n2] {
            n1++
        } else {
            n2++
        }
    }
    
    return out
}