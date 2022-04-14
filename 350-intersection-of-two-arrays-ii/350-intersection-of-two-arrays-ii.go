
// appraoch : freq map
func intersect(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    out := []int{}
    if len(nums1) < len(nums2) {
        for _, ele := range nums1 {
            m[ele]++
            
        }
        for i := 0; i < len(nums2); i++ {
            if _, ok := m[nums2[i]]; ok {
                m[nums2[i]]--
                if freq := m[nums2[i]]; freq == 0 {
                    delete(m, nums2[i])
                }
                out = append(out, nums2[i])
            }
        }
    } else {
        for _, ele := range nums2 {
            m[ele]++
        }
        for i := 0; i < len(nums1); i++ {
            if _, ok := m[nums1[i]]; ok {
                m[nums1[i]]--
                if freq := m[nums1[i]]; freq == 0 {
                    delete(m, nums1[i])
                }
                out = append(out, nums1[i])
            }
        }
    }
    return out
    
}