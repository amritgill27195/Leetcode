
// appraoch : freq map
// func intersect(nums1 []int, nums2 []int) []int {
//     m := map[int]int{}
//     out := []int{}
//     if len(nums1) < len(nums2) {
//         for _, ele := range nums1 {
//             m[ele]++
            
//         }
//         for i := 0; i < len(nums2); i++ {
//             if _, ok := m[nums2[i]]; ok {
//                 m[nums2[i]]--
//                 if freq := m[nums2[i]]; freq == 0 {
//                     delete(m, nums2[i])
//                 }
//                 out = append(out, nums2[i])
//             }
//         }
//     } else {
//         for _, ele := range nums2 {
//             m[ele]++
//         }
//         for i := 0; i < len(nums1); i++ {
//             if _, ok := m[nums1[i]]; ok {
//                 m[nums1[i]]--
//                 if freq := m[nums1[i]]; freq == 0 {
//                     delete(m, nums1[i])
//                 }
//                 out = append(out, nums1[i])
//             }
//         }
//     }
//     return out
// }

// approach sort + two pointers
// time: 
//      o(n1logn1) + o(n2logn2) +  o(max(nums1,nums2)) -- since we sort both
// space: o(1) if we do not count output array as part of space
func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    n1 := 0
    n2 := 0
    
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