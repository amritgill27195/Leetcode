func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {
        return intersect(nums2, nums1)
    }
    
    nums1FreqMap := map[int]int{}
    for _, ele := range nums1 {
        nums1FreqMap[ele]++
    }
    
    out := []int{}
    for _, ele := range nums2 {
        val, ok := nums1FreqMap[ele]
        if ok && val != 0 {
            nums1FreqMap[ele]--
            out = append(out, ele)
            if val == 1 {
                delete(nums1FreqMap, ele)
            }
        }
    }
    return out
    
}